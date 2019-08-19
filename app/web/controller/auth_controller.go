package controller

import (
	"fmt"
	"komodo/app/dto/request"
	"komodo/app/dto/response"
	"komodo/app/web/models"
	"komodo/app/web/utils"
	"komodo/database"
	"log"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var FormRegister request.Register
	var user models.User
	var validationError []string

	if c.BindJSON(&FormRegister) == nil {
		db := database.Connect()
		row := db.QueryRow("SELECT email FROM users WHERE email = ? ;", FormRegister.Email).Scan(&user.Email)
		if row == nil {
			validationError = append(validationError, "Email already Registered")
		}
		if len(validationError) > 0 {
			c.JSON(response.BAD_REQUEST, gin.H{
				"message": validationError,
			})
		} else {
			uuid, _ := utils.NewUUID()
			password := utils.HashAndSalt([]byte(FormRegister.Password))
			stmt, err := db.Prepare("INSERT INTO users (uuid,email,password) values(?,?,?);")
			if err != nil {
				log.Print(err)
			}
			_, err = stmt.Exec(uuid, FormRegister.Email, password)
			if err != nil {
				fmt.Print(err.Error())
				return
			}
			defer stmt.Close()
			to := FormRegister.Email
			_ = utils.SendMail(to, "Confirmation Email", uuid)
			c.JSON(response.CREATED, gin.H{
				"message": response.CREATED_MESSAGE,
			})
		}

	}
	return
}

func ConfirmAccount(c *gin.Context) {
	db := database.Connect()
	uuid := c.Query("code")
	row, err := db.Prepare("UPDATE users set confrim= 1 WHERE uuid= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = row.Exec(uuid)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		c.JSON(response.OK, gin.H{
			"message": "Your account active",
		})
	}
	return
}
