package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv" // Import godotenv
	"github.com/stfuxbm/minorum/internal/database"
	"github.com/stfuxbm/minorum/internal/routes"
)

func main() {
	// Memuat file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Koneksi ke MongoDB
	database.MongoConnect()

	// Setup semua route
	mux := routes.SetupRoutes()

	// Ambil PORT dari environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback ke 8080 jika tidak ada variabel PORT
	}

	// Tampilkan log server jalan
	log.Println("Server started at http://localhost:" + port)

	// Jalankan server di port yang sesuai
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
