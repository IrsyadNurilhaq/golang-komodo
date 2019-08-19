package utils

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserId(c *gin.Context) int {
	var user_id int
	authHeader := c.Request.Header.Get("Authorization")
	arrayAuthHeader := strings.Split(authHeader, " ")
	if len(arrayAuthHeader) == 2 {
		if arrayAuthHeader[0] == "Bearer" {
			tokenString := arrayAuthHeader[1]
			claims := jwt.MapClaims{}
			_, _ = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if jwt.GetSigningMethod("HS256") != token.Method {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte("secret"), nil
			})
			for _, val := range claims {
				user_id = val.(int)
			}
		}
	}
	return user_id
}
