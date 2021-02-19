package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 参数接收格式为json
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	r.Run(":8080")
}
