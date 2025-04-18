package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/stfuxbm/minorum/internal/database"
	models "github.com/stfuxbm/minorum/internal/models/ordos"
)

// DecodeOrdo mengubah request body (JSON) menjadi struct Ordo
func DecodeOrdo(r *http.Request) (models.Ordo, error) { // Mengembalikan tipe models.Ordo
	var Ordo models.Ordo
	if err := json.NewDecoder(r.Body).Decode(&Ordo); err != nil {
		return Ordo, fmt.Errorf("invalid JSON format") // Menggunakan string pesan kesalahan yang lebih jelas
	}
	return Ordo, nil
}

// ValidateOrdo memvalidasi data Ordo:

func ValidateOrdo(Ordo models.Ordo) error {
	if Ordo.OrderName == "" || Ordo.LatinName == "" || Ordo.Nickname == "" || Ordo.CurrentLeader == "" {
		return fmt.Errorf("Ordo and author name are required") // Menggunakan pesan kesalahan yang lebih jelas
	}
	return nil
}

// CheckDuplicateOrdo mengecek apakah Ordo dengan nama yang sama sudah ada di database
func CheckDuplicateOrdo(r *http.Request, Ordo models.Ordo) error {
	collection := database.DB.Collection("ordo")

	filter := map[string]interface{}{
		"order_name": Ordo.OrderName, // Perbaikan: gunakan key yang sesuai dengan field JSON/DB
	}

	count, err := collection.CountDocuments(r.Context(), filter)
	if err != nil {
		return fmt.Errorf("internal server error")
	}
	if count > 0 {
		return fmt.Errorf("Ordo already exists")
	}
	return nil
}

// SaveOrdo menyimpan data Ordo ke MongoDB, sekaligus set CreatedAt dan UpdatedAt
func SaveOrdo(r *http.Request, Ordo *models.Ordo) error {
	collection := database.DB.Collection("ordo")
	now := time.Now().UTC()

	// Jika data sudah ada, perbarui UpdatedAt
	if Ordo.CreatedAt.IsZero() {
		Ordo.CreatedAt = now // Jika pertama kali dibuat, set CreatedAt
	}
	Ordo.UpdatedAt = now // Selalu update UpdatedAt

	// Menyimpan atau mengupdate data ke database
	_, err := collection.InsertOne(r.Context(), Ordo)
	if err != nil {
		return fmt.Errorf("internal server error") // Pesan kesalahan yang lebih jelas
	}
	return nil
}
