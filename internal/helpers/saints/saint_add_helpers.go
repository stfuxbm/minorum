package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/stfuxbm/minorum/internal/database"
	models "github.com/stfuxbm/minorum/internal/models/saints" // ganti path model sesuai
)

// DecodeSaint mengubah request body (JSON) menjadi struct Saints
func DecodeSaint(r *http.Request) (models.Saints, error) {
	var saint models.Saints
	if err := json.NewDecoder(r.Body).Decode(&saint); err != nil {
		return saint, fmt.Errorf("Invalid JSON format")
	}
	return saint, nil
}

// Validasi Saint
func ValidateSaint(saint models.Saints) error {
	if saint.Name == "" || saint.LatinName == "" || saint.NickName == "" {
		return fmt.Errorf("Name, Latin Name, and Nickname are required")
	}
	return nil
}

// Check Duplicate Saint
func CheckDuplicateSaint(r *http.Request, saint models.Saints) error {
	collections := database.DB.Collection("saint") // ganti nama koleksi sesuai

	filter := map[string]interface{}{
		"name": saint.Name,
	}
	count, err := collections.CountDocuments(r.Context(), filter)
	if err != nil {
		return fmt.Errorf("Internal server error")
	}
	if count > 0 {
		return fmt.Errorf("Saint already exists")
	}
	return nil
}

// Save Saint
func SaveSaint(r *http.Request, saint *models.Saints) error {
	collections := database.DB.Collection("saint") // ganti nama koleksi sesuai

	now := time.Now().UTC()

	if saint.CreatedAt.IsZero() {
		saint.CreatedAt = now
	}
	saint.UpdatedAt = now

	// update data ke database
	_, err := collections.InsertOne(r.Context(), saint)
	if err != nil {
		return fmt.Errorf("Internal server error")
	}
	return nil
}
