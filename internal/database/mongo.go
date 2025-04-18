package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/stfuxbm/minorum/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database // Variabel global untuk menyimpan koneksi database

func MongoConnect() {
	// Ambil URI dan nama database dari file .env melalui config package
	uri := config.GetMongoURI()
	dbName := config.GetMongoDBName()

	// Buat context dengan timeout 10 detik untuk koneksi ke MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Pastikan context dibatalkan setelah selesai (menghindari memory leak)

	// Inisialisasi koneksi ke MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		// Jika koneksi gagal, hentikan aplikasi dengan pesan error
		log.Fatal("Failed to connect to MongoDB:", err)
		return
	}

	// Cek apakah koneksi berhasil (ping ke server MongoDB)
	err = client.Ping(ctx, nil)
	if err != nil {
		// Jika tidak bisa ping, hentikan aplikasi juga
		log.Fatal("Unable to ping MongoDB:", err)
	}

	// Simpan koneksi database ke variabel global DB
	DB = client.Database(dbName)

	// Tampilkan pesan sukses ke terminal
	fmt.Println("Successfully connected to MongoDB")
}
