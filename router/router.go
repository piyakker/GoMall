package router

import (
	"context"
	"net/http"
	"time"

	config "github.com/piyakker/GoMall/configs"

	"github.com/piyakker/GoMall/health"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
		defer cancel()

		checkers := []health.Checker{
			health.NewPostgresChecker(cfg),
			health.NewRedisChecker(cfg),
			health.NewRabbitMQChecker(cfg),
		}

		if err := health.RunAll(ctx, checkers...); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "unhealthy",
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	return r
}
