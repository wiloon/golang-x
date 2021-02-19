package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type wechatInvoke struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

func main() {
	router := gin.Default()

	// router.Use(limits.RequestSizeLimiter(10))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/Services/Service.svc/GetCsCode", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/ping", func(c *gin.Context) {
		wi := &wechatInvoke{}
		if c.BindJSON(wi) == nil {
			fmt.Printf("code: %v\n", wi.Code)
		}

		if len(c.Errors) > 0 {
			fmt.Println(c.Errors)
			return
		}
		c.JSON(200, gin.H{
			"jwt_token": "jwt_token0",
		})
	})

	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Println(fmt.Sprintf("id: %v, page: %v, name: %v, msg: %v", id, page, name, message))
		if len(c.Errors) > 0 {
			fmt.Printf("errors: %v", c.Errors)
			return
		}
		c.JSON(200, gin.H{
			"jwt_token": "jwt_token0",
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
