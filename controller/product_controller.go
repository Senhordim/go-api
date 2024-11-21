package controller

import (
	"net/http"

	"github.com/Senhordim/go-api/model"
	"github.com/gin-gonic/gin"
)

type productController struct{}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:          1,
			Name:        "Laptop",
			Description: "Lenovo ThinkPad",
			Price:       500.50,
		},
		{
			ID:          2,
			Name:        "Nitendo Switch",
			Description: "Nitendo Switch Console",
			Price:       1500.81,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
