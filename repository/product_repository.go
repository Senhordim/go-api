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

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, error := pr.connecion.Prepare("INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id")

	if error != nil {
		fmt.Println(error)
		return 0, error
	}

	error = query.QueryRow(
		product.Name,
		product.Description,
		product.Price,
	).Scan(&id)

	if error != nil {
		fmt.Println(error)
		return 0, error
	}

	defer query.Close()

	return id, nil
}
