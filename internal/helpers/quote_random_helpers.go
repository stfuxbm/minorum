// helpers/get_all_quotes.go
package helpers

import (
	"context"
	"math/rand"

	"github.com/stfuxbm/quote-saints/internal/database"
	"github.com/stfuxbm/quote-saints/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetRandomQuotes mengambil semua quotes dari MongoDB
func GetRandomQuotes(ctx context.Context) ([]models.Quote, error) {
	// Akses koleksi quotes pada database MongoDB
	quotesCollection := database.DB.Collection("quotes")

	// Ambil semua quotes
	cursor, err := quotesCollection.Find(ctx, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Menyimpan quotes dalam slice
	var quotes []models.Quote
	for cursor.Next(ctx) {
		var quote models.Quote
		if err := cursor.Decode(&quote); err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}

	// Jika tidak ada quotes, kembalikan error
	if len(quotes) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	// Pilih quote acak dari slice
	randomIndex := rand.Intn(len(quotes)) // Ambil index acak
	randomQuote := quotes[randomIndex]    // Pilih quote sesuai index

	return []models.Quote{randomQuote}, nil // Mengembalikan slice berisi 1 quote acak
}
