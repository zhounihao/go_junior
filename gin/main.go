package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, go")
	})
	server.Run("0.0.0.0:8080") // 默认监听 0.0.0.0:8080
}
