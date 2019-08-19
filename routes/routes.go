package routes

import (
	controller "komodo/app/web/controller"
	Auth "komodo/app/web/middleware"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/login", controller.LoginHandler)
		v1.POST("/register", controller.Register)
		v1.GET("/confirmation", controller.ConfirmAccount)
	}
	v1.Use(Auth.Auth)
	{
		v1.GET("/tweets", controller.GetLastTweet)
		v1.POST("/follow", controller.FollowUser)
		v1.GET("/followers", controller.GetFollower)
		v1.GET("/following", controller.GetFollowing)
	}

	return router
}
