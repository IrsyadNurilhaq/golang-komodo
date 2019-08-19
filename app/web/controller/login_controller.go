package controller

import (
	"komodo/app/dto/response"
	"komodo/app/web/models"
	"komodo/app/web/utils"
	"komodo/database"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var user models.User

	err := c.Bind(&user)
	if err != nil {
		c.JSON(response.BAD_REQUEST, gin.H{
			"status":  response.BAD_REQUEST,
			"message": "JSON Format Only",
		})
		return
	}
	plainPwd := user.Password
	db := database.Connect()
	row := db.QueryRow("SELECT password FROM users WHERE email = ? ;", user.Email).Scan(&user.Password)
	if row != nil {
		c.JSON(response.UN_AUTHORIZED, gin.H{
			"message": "Your email not registered",
		})
		return
	}
	email := user.Email
	pwd := utils.ComparePasswords(user.Password, []byte(plainPwd))
	if pwd == true {
		row = db.QueryRow("SELECT confrim,id FROM users WHERE email = ? AND password = ? ", email, user.Password).Scan(&user.Confrim, &user.Id)
		if user.Confrim == false {
			c.JSON(response.UN_AUTHORIZED, gin.H{
				"message": "Your not Active",
			})
			return
		} else {
			sign := jwt.New(jwt.GetSigningMethod("HS256"))
			claims := sign.Claims.(jwt.MapClaims)
			claims["user_id"] = user.Id
			token, err := sign.SignedString([]byte("secret"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				c.Abort()
			}
			c.JSON(response.OK, gin.H{
				"status": response.OK_MESSAGE,
				"token":  token,
			})
		}

	} else {
		c.JSON(response.UN_AUTHORIZED, gin.H{
			"message": "Your password is wrong",
		})
		return
	}
	return
}
