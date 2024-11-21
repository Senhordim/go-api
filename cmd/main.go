package main

import (
	"github.com/Senhordim/go-api/config/db"
	"github.com/Senhordim/go-api/controller"
	"github.com/Senhordim/go-api/repository"
	"github.com/Senhordim/go-api/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin
	server := gin.Default()

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	// Repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// UseCase
	ProductUseCase := usecase.NewProductUseCase(
		ProductRepository,
	)

	// Controller
	ProductController := controller.NewProductController(
		ProductUseCase,
	)

	// Routes
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello Go Lang",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	// Run Server
	server.Run(":3000")
}
