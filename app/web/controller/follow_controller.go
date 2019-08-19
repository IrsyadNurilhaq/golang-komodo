package controller

import (
	"fmt"
	"komodo/app/dto/request"
	"komodo/app/dto/response"
	"komodo/app/web/models"
	"komodo/database"
	"log"

	"github.com/gin-gonic/gin"
)

func FollowUser(c *gin.Context) {
	var FormFollow request.Follow
	var user models.User
	var follow models.Follow
	session, _ := store.Get(c.Request, "credential")
	userId := session.Values["user_id"]

	if c.BindJSON(&FormFollow) == nil {
		db := database.Connect()
		row := db.QueryRow("SELECT id FROM users WHERE email = ? ;", FormFollow.Email).Scan(&user.Id)
		if row == nil {
			stmt := db.QueryRow("SELECT id FROM follows WHERE user_id = ? AND `to` = ?", userId, user.Id).Scan(&follow.Id)
			if stmt == nil {
				c.JSON(response.BAD_REQUEST, gin.H{
					"message": "Already follow",
				})
			} else {
				rowSql, err := db.Prepare("INSERT INTO `follows` (`user_id`, `to`) VALUES (?, ?)")
				if err != nil {
					log.Print(err)
				}
				_, err = rowSql.Exec(userId, user.Id)
				if err != nil {
					fmt.Print(err.Error())
					return
				}

				c.JSON(response.OK, gin.H{
					"message": "success",
				})
			}

		} else {
			c.JSON(response.BAD_REQUEST, gin.H{
				"message": "User undifined",
			})
		}

	}
	return
}

func GetFollower(c *gin.Context) {
	var (
		follower  response.FollowerFollowing
		followers []response.FollowerFollowing
	)

	session, _ := store.Get(c.Request, "credential")
	userId := session.Values["user_id"]

	db := database.Connect()
	rows, err := db.Query("SELECT users.email FROM follows LEFT JOIN users ON follows.user_id = users.id WHERE `to` = ?", userId)
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&follower.Email)
		followers = append(followers, follower)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(response.OK, gin.H{
		"data": followers,
	})
	return
}

func GetFollowing(c *gin.Context) {
	var (
		follower  response.FollowerFollowing
		followers []response.FollowerFollowing
	)

	session, _ := store.Get(c.Request, "credential")
	userId := session.Values["user_id"]

	db := database.Connect()
	rows, err := db.Query("SELECT users.email FROM follows LEFT JOIN users ON follows.to = users.id WHERE user_id = ?", userId)
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&follower.Email)
		followers = append(followers, follower)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(response.OK, gin.H{
		"data": followers,
	})
	return
}
