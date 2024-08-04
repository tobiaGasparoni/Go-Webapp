package products

import (
	"fmt"
	"log"
	"os"

	"github.com/tobiaGasparoni/Go-Webapp/pkg/db/utils"
	"github.com/tobiaGasparoni/Go-Webapp/pkg/models"
)

// GetAllProducts retrieves all rows from the products table in the db and converts them into the adecuate entities
func GetAllProducts() ([]models.Product, error) {
	conn, err := utils.CreateDbConnection()
	if err != nil {
		os.Exit(1)
	}

	rows, err := DbQueryGetAllProducts(conn)
	if err != nil {
		return nil, err
	}

	products := []models.Product{}

	var name, category string
	var id, price int

	for rows.Next() {
		err := rows.Scan(&id, &name, &price, &category)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		products = append(products, models.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Category: category,
		})
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows!", err)
	}

	fmt.Println("----------------------------")

	return products, nil
}

// CreateProduct saves a new product in the db
func CreateProduct(product models.InputProduct) (models.Product, error) {
	conn, err := utils.CreateDbConnection()
	if err != nil {
		os.Exit(1)
	}
	log.Println(conn)

	return models.Product{
		Name:     product.Name,
		Category: product.Category,
		Price:    product.Price,
	}, nil
}
