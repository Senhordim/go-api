package main

import (
	"github.com/Senhordim/go-api/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	ProductController := controller.NewProductController()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello Go Lang",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":3000")
}
