package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/stfuxbm/quote-saints/internal/database"
	models "github.com/stfuxbm/quote-saints/internal/models/quotes"
)

// DecodeQuote mengubah request body (JSON) menjadi struct Quote
func DecodeQuote(r *http.Request) (models.Quote, error) { // Mengembalikan tipe models.Quote
	var quote models.Quote
	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		return quote, fmt.Errorf("invalid JSON format") // Menggunakan string pesan kesalahan yang lebih jelas
	}
	return quote, nil
}

// ValidateQuote memvalidasi data quote:
// - quote dan author name tidak boleh kosong
func ValidateQuote(quote models.Quote) error {
	if quote.Quote == "" || quote.Author.Name == "" {
		return fmt.Errorf("quote and author name are required") // Menggunakan pesan kesalahan yang lebih jelas
	}
	return nil
}

// CheckDuplicateQuote mengecek apakah quote dengan isi dan nama author yang sama sudah ada di database
func CheckDuplicateQuote(r *http.Request, quote models.Quote) error {
	collection := database.DB.Collection("quotes")
	filter := map[string]interface{}{
		"quote":       quote.Quote,
		"author.name": quote.Author.Name,
	}
	count, err := collection.CountDocuments(r.Context(), filter)
	if err != nil {
		return fmt.Errorf("internal server error") // Pesan kesalahan yang lebih jelas
	}
	if count > 0 {
		return fmt.Errorf("quote already exists") // Pesan kesalahan yang lebih jelas
	}
	return nil
}

// SaveQuote menyimpan data quote ke MongoDB, sekaligus set CreatedAt dan UpdatedAt
func SaveQuote(r *http.Request, quote *models.Quote) error {
	collection := database.DB.Collection("quotes")
	now := time.Now().UTC()

	// Jika data sudah ada, perbarui UpdatedAt
	if quote.CreatedAt.IsZero() {
		quote.CreatedAt = now // Jika pertama kali dibuat, set CreatedAt
	}
	quote.UpdatedAt = now // Selalu update UpdatedAt

	// Menyimpan atau mengupdate data ke database
	_, err := collection.InsertOne(r.Context(), quote)
	if err != nil {
		return fmt.Errorf("internal server error") // Pesan kesalahan yang lebih jelas
	}
	return nil
}
