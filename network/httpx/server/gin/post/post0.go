package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LOGIN struct {
	code     string `json:"code" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.String(200, "on")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	r.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.POST("/ping", func(c *gin.Context) {
		var login LOGIN
		c.BindJSON(&login)
		c.JSON(200, gin.H{"status": login.code})
	})
	r.Run(":8081") // listen an
}
