package router

import (
	"github.com/DeveloperGerald/TurtleSoup/controller"
	"github.com/DeveloperGerald/TurtleSoup/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	// 配置 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // 允许的来源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		AllowCredentials: true,                                                // 是否允许携带凭证
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * 3600, // 预检请求的缓存时间（秒）
	}))

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

	chat := v1.Group("/game-master", middleware.AuthorizationMiddleware())
	chat.POST("/answer", controller.GiveAnswerController)
	chat.GET("/random-story", controller.GetRandomStoryController)
}
