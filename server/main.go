package main

import (
	"github.com/gin-gonic/gin"
	"server/handlers"
)


func main() {
	r := gin.Default()
	r.POST("/header", handlers.PostHeader)
	r.POST("/footer", handlers.PostFooter)
	r.POST("/card", handlers.PostCard)
	r.Run("0.0.0.0:5000")
}
