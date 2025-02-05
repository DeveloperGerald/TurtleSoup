package router

import (
	"github.com/DeveloperGerald/TurtleSoup/controller"
	"github.com/DeveloperGerald/TurtleSoup/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/ttsoup/api/v1")

	user := v1.Group("/user")
	user.POST("/register", controller.RegisterUserController)
	user.POST("/login", controller.LoginUserController)

	story := v1.Group("/story", middleware.AuthorizationMiddleware())
	story.GET("")
	story.GET("/:id")
	story.POST("", controller.CreateStoryController)
	story.PUT("/:id")
	story.DELETE("/:id")
}
