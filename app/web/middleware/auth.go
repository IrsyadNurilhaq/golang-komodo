package middleware

import (
	"fmt"
	"komodo/app/dto/response"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func Auth(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader != "" {
		arrayAuthHeader := strings.Split(authHeader, " ")
		if len(arrayAuthHeader) == 2 {
			if arrayAuthHeader[0] == "Bearer" {
				tokenString := arrayAuthHeader[1]
				claims := jwt.MapClaims{}
				token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
					if jwt.GetSigningMethod("HS256") != token.Method {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}
					return []byte("secret"), nil
				})

				if token != nil && err == nil {
					session, _ := store.Get(c.Request, "credential")
					session.Values["user_id"] = claims["user_id"]
					session.Save(c.Request, c.Writer)
					c.Next()
				}
			} else {
				c.JSON(response.UN_AUTHORIZED, gin.H{
					"message": "Invalid token format",
				})
				c.Abort()
			}
		} else {
			c.JSON(response.UN_AUTHORIZED, gin.H{
				"message": "Invalid token format",
			})
			c.Abort()
		}
	} else {
		c.JSON(response.UN_AUTHORIZED, gin.H{
			"message": "You can't access this url",
		})
		c.Abort()
	}
}
