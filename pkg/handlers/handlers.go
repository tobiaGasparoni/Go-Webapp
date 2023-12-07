package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tobiaGasparoni/Go-Webapp/pkg/config"
	"github.com/tobiaGasparoni/Go-Webapp/pkg/models"

	"github.com/jackc/pgx/v5"
)

// Repo the respository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Products(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	conn, err := pgx.Connect(context.Background(), "postgres://citizix_user:S3cret@postgres:5432/citizix_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	log.Println("Connected to database!")

	err, products := getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	jsonResp, err := json.Marshal(products)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func getAllRows(conn *pgx.Conn) (error, []models.Product) {
	rows, err := conn.Query(context.Background(), "SELECT id, name, price, category FROM products")
	if err != nil {
		log.Println(err)
		return err, nil
	}
	defer rows.Close()

	products := []models.Product{}

	var name, category string
	var id, price int

	for rows.Next() {
		err := rows.Scan(&id, &name, &price, &category)
		if err != nil {
			log.Println(err)
			return err, nil
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

	return nil, products
}

func (m *Repository) ProductDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := models.Product{
		Name:     "Laptop",
		Price:    3411,
		Category: "Electronics",
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
