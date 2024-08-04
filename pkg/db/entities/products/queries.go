package products

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

// DbQueryGetAllProducts uses the created db connection to query all the products in the db
func DbQueryGetAllProducts(conn *pgx.Conn) (pgx.Rows, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, name, price, category FROM products")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(rows)
	defer rows.Close()
	defer conn.Close(context.Background())

	return rows, nil
}
