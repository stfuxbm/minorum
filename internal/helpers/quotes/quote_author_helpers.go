package helpers

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/stfuxbm/minorum/internal/database"
	models "github.com/stfuxbm/minorum/internal/models/quotes"
)

// GetQuotesBySaintName mencari kutipan berdasarkan nama santo dari database
func GetQuotesBySaintName(ctx context.Context, order string) ([]models.Quote, error) {
	// Ambil koleksi quotes dari database
	collection := database.DB.Collection("quote")

	// Membuat filter pencarian untuk 'author' dan 'category' dengan pencarian case-insensitive
	filter := bson.M{
		"$or": []bson.M{
			{
				"author": bson.M{ // Pencarian case-insensitive pada field 'author'
					"$regex":   order, // Pencarian case-insensitive
					"$options": "i",   // Tambahkan opsi "i" untuk case-insensitive
				},
			},
			{
				"category": bson.M{ // Pencarian case-insensitive pada field 'category'
					"$regex":   order, // Pencarian case-insensitive
					"$options": "i",   // Tambahkan opsi "i" untuk case-insensitive
				},
			},
		},
	}

	// Ambil quotes berdasarkan filter
	var quotes []models.Quote
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Handle the case where no documents are found
			return nil, nil // or return an empty slice
		}
		return nil, err
	}
	defer cursor.Close(ctx)

	// Parse hasil cursor ke dalam slice quotes
	if err := cursor.All(ctx, &quotes); err != nil {
		return nil, err
	}

	return quotes, nil
}
