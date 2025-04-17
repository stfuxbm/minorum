package helpers

import (
	"context"

	"github.com/stfuxbm/quote-saints/internal/database"
	models "github.com/stfuxbm/quote-saints/internal/models/quotes"
	"go.mongodb.org/mongo-driver/bson"
)

// GetQuotesByCategory mencari kutipan berdasarkan kategori dari database
func GetQuotesByCategories(ctx context.Context, categories []string) ([]models.Quote, error) {
	collection := database.DB.Collection("quotes")

	filter := bson.M{
		"category": bson.M{
			"$in": categories, // cocokkan dengan salah satu dari daftar kategori
		},
	}

	var quotes []models.Quote
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &quotes); err != nil {
		return nil, err
	}

	return quotes, nil
}
