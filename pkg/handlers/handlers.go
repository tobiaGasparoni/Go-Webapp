package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tobiaGasparoni/Go-Webapp/pkg/config"
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

func (m *Repository) Products(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := []models.Product{}
	resp = append(resp, models.Product{
		Name:     "Laptop",
		Price:    3411,
		Category: "Electronics",
	})
	resp = append(resp, models.Product{
		Name:     "SmartPhone",
		Price:    343,
		Category: "Electronics",
	})
	resp = append(resp, models.Product{
		Name:     "The Lord of the Rings",
		Price:    343,
		Category: "Books",
	})

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
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
