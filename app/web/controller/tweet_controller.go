package controller

import (
	"fmt"
	"komodo/app/dto/response"
	"komodo/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func GetLastTweet(c *gin.Context) {
	var (
		tweet  response.Tweet
		tweets []response.Tweet
	)

	session, _ := store.Get(c.Request, "credential")
	userId := session.Values["user_id"]

	offset := c.Query("offset")
	db := database.Connect()

	rows, err := db.Query(`SELECT tweets.id,users.email, tweets.tweet FROM tweets INNER JOIN follows ON tweets.user_id = follows.to 
	LEFT JOIN users ON tweets.user_id = users.id WHERE follows.user_id = ? LIMIT 10 OFFSET ?`, userId, offset)
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&tweet.Id, &tweet.Email, &tweet.Tweet)
		tweets = append(tweets, tweet)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(response.OK, gin.H{
		"data": tweets,
	})
	return

}
