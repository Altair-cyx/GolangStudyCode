package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向

func main() {
	r := gin.Default()

	r.GET("/redirect", func(c *gin.Context) {
		// 支持内部和外部的重定向
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

	r.Run(":8080")
}
