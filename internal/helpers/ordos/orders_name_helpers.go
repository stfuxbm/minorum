package helpers

import (
	"context"

	"github.com/stfuxbm/minorum/internal/database"
	models "github.com/stfuxbm/minorum/internal/models/ordos"
	"go.mongodb.org/mongo-driver/bson"
)

// GetOrdoByOrderName mencari ordo berdasarkan nama ordo atau nickname dari database dengan pencarian case-insensitive
func GetOrdoByOrderName(ctx context.Context, order string) (*models.Ordo, error) {
	// Ambil koleksi orders dari database
	collection := database.DB.Collection("ordo")

	// Cari ordo yang mengandung nama ordo atau nickname dalam field 'order_name' atau 'nickname' (case-insensitive)
	filter := bson.M{
		"$or": []bson.M{
			{
				"order_name": bson.M{
					"$regex":   order, // Pencarian case-insensitive
					"$options": "i",   // Tambahkan opsi "i" untuk case-insensitive
				},
			},
			{
				"nickname": bson.M{
					"$regex":   order, // Pencarian case-insensitive
					"$options": "i",   // Tambahkan opsi "i" untuk case-insensitive
				},
			},
			{
				"title": bson.M{
					"$regex":   order, // Pencarian case-insensitive
					"$options": "i",   // Tambahkan opsi "i" untuk case-insensitive
				},
			},
			{
				"category": bson.M{
					"$regex":   order, // Pencarian case-insensitive
					"$options": "i",   // Tambahkan opsi "i" untuk case-insensitive
				},
			},
			{
				"founder_name": bson.M{
					"$regex":   order, // Pencarian case-insensitive
					"$options": "i",   // Tambahkan opsi "i" untuk case-insensitive
				},
			},
		},
	}

	// Ambil ordo berdasarkan filter
	var ordo models.Ordo
	err := collection.FindOne(ctx, filter).Decode(&ordo)
	if err != nil {
		return nil, err
	}

	return &ordo, nil
}
