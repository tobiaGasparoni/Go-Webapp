package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tobiaGasparoni/Go-Webapp/pkg/config"
	"github.com/tobiaGasparoni/Go-Webapp/pkg/db/entities/products"
	"github.com/tobiaGasparoni/Go-Webapp/pkg/models"
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

// GetAllProducts returns all the products in the db
func (m *Repository) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := products.GetAllProducts()
	if err != nil {
		log.Fatal(err)
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}

// CreateProduct saves a new product in the db
func (m *Repository) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.InputProduct

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	log.Println(product)

	resp, err := products.CreateProduct(product)
	if err != nil {
		log.Fatal(err)
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}
