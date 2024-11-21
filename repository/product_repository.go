package repository

import (
	"database/sql"
	"fmt"

	"github.com/Senhordim/go-api/model"
)

type ProductRepository struct {
	connecion *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connecion: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, name, description, price FROM products"
	rows, err := pr.connecion.Query(query)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, nil
	}

	var productsList []model.Product
	var product model.Product

	for rows.Next() {
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, nil
		}

		productsList = append(productsList, product)

	}

	return productsList, nil
}
