package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	//fmt.Println("123")
	r := gin.Default()
	//监听并返回信息
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// AsciiJSON
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})
	// HTML渲染
	//r.LoadHTMLGlob("templates/*")
	//r.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	//		"title": "Main website",
	//	})
	//})
	//使用不同目录下名称相同的模板
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts Index",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users Index",
		})
	})
	// 自定义模板渲染器
	//html := template.Must(template.ParseFiles("file1", "file2"))
	//r.SetHTMLTemplate(html)

	//自定义分隔符
	r.Delims("{[{", "}]}") // 默认为 {{}}
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLFiles("./templates/raw.tmpl")
	r.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2021, 02, 19, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":8080") //监听127.0.0.1:8080启动服务
	//r := setupRouter()
	//r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}
