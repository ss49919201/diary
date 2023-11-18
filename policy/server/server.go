package server

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/ss49919201/diary/policy/controller"
)

func Run() error {
	gin.SetMode("release")

	ginEngine := gin.New()
	ginEngine.Use(logger.SetLogger())

	ginEngine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	ginEngine.POST("/policy", controller.CreatePolicy)

	return ginEngine.Run("11100")
}
