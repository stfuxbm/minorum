package helpers

import (
	"context"

	"github.com/stfuxbm/quote-saints/internal/database"
	models "github.com/stfuxbm/quote-saints/internal/models/quotes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetQuotesBySaintName mencari kutipan berdasarkan nama santo dari database
func GetQuotesBySaintName(ctx context.Context, name string) ([]models.Quote, error) {
	// Ambil koleksi quotes dari database
	collection := database.DB.Collection("quotes")

	// Cari kutipan yang mengandung nama santo dalam field 'author.name' (menggunakan pencarian case-insensitive)
	filter := bson.M{
		"author.name": bson.M{
			"$regex":   name,
			"$options": "i", // Case insensitive
		},
	}

	// Ambil quotes berdasarkan filter
	var quotes []models.Quote
	cursor, err := collection.Find(ctx, filter, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Parse hasil cursor ke dalam slice quotes
	if err := cursor.All(ctx, &quotes); err != nil {
		return nil, err
	}

	return quotes, nil
}
