package helpers

import (
	"context"

	"github.com/stfuxbm/minorum/internal/database"
	models "github.com/stfuxbm/minorum/internal/models/saints" // ganti path model sesuai
	"go.mongodb.org/mongo-driver/bson"
)

// GetSaintByNameOrNickname mencari saint berdasarkan name atau nickname dari database dengan pencarian case-insensitive
func GetSaintByNameOrNickname(ctx context.Context, order string) (*models.Saints, error) {
	// Ambil koleksi saints dari database
	collection := database.DB.Collection("saint") // ganti nama koleksi sesuai

	// Buat filter pencarian dengan regex (case-insensitive) untuk name atau nick_name
	filter := bson.M{
		"$or": []bson.M{
			{
				"name": bson.M{
					"$regex":   order,
					"$options": "i",
				},
			},
			{
				"nick_name": bson.M{
					"$regex":   order,
					"$options": "i",
				},
			},
			{
				"latin_name": bson.M{
					"$regex":   order,
					"$options": "i",
				},
			},
			{
				"patronage": bson.M{
					"$regex":   order,
					"$options": "i",
				},
			},
			{
				"symbols": bson.M{
					"$regex":   order,
					"$options": "i",
				},
			},
		},
	}

	// Eksekusi query
	var saint models.Saints
	err := collection.FindOne(ctx, filter).Decode(&saint)
	if err != nil {
		return nil, err
	}

	return &saint, nil
}
