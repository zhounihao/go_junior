package main

import (
	"github.com/gin-gonic/gin"
	"go-project_junior/webhook/internal/web"
)

func main() {
	server := gin.Default()
	u := web.NewUserHandler()
	u.RegisterRouters(server)
	server.Run(":8080")
}
