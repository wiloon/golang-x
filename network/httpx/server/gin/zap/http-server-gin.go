package main

import (
	"github.com/rcrowley/go-metrics"

	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

func main() {
	go metrics.Log(
		metrics.DefaultRegistry,
		10*time.Second,
		log.New(os.Stderr, "metrics: ", log.Lmicroseconds),
	)

	r := gin.New()
	logger, _ := zap.NewProduction()

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
		meter := metrics.GetOrRegisterMeter("foo", nil)
		meter.Mark(1)
	})
	_ = r.Run(":8080")
}
