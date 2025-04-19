package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/stfuxbm/minorum/internal/database"
	"github.com/stfuxbm/minorum/internal/routes"
)

func main() {
	// Coba load .env untuk local development
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
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
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
