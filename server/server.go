package server

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ss49919201/diary/customlogger"
	"github.com/ss49919201/diary/policy/controller"
)

func Run(logger customlogger.Logger) error {
	gin.SetMode("release")

	ginEngine := gin.New()
	ginEngine.Use(func(ctx *gin.Context) {
		now := time.Now()
		ctx.Next()
		latency := time.Since(now)
		logger.Info("access log",
			"status", ctx.Writer.Status(),
			"method", ctx.Request.Method,
			"path", ctx.Request.URL.Path,
			"ip", ctx.ClientIP(),
			"ua", ctx.Request.UserAgent(),
			"latency", fmt.Sprintf("%fms", float64(latency)/float64(time.Millisecond)),
		)
	})

	ginEngine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	ginEngine.POST("/policy", controller.CreatePolicy)

	return ginEngine.Run("localhost:11100")
}
