package main

import (
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:user binding:"required"`
	Password string `form:password binding:"required"`
}

func main() {
	r := gin.Default()
	// 参数接收格式为json
	r.POST("/login", func(c *gin.Context) {
		// 可以显示绑定声明绑定multipart form:
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定
		var form LoginForm
		// 在这种情况下, 将自动选择合适的绑定
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	r.Run(":8080")
}
