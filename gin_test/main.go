package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	ua := ""

	// ミドルウェアを使用
	engine.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})
	engine.LoadHTMLGlob("templates/*")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "hello gin",
		})
	})
	engine.Run(":3000")
}
