package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/stfuxbm/minorum/config"
	"github.com/stfuxbm/minorum/internal/database"
	"github.com/stfuxbm/minorum/internal/routes"
)

func main() {
	// Load environment variable dari file .env (gunakan konfigurasi Render untuk ENV di production)
	config.LoadEnv()

	// Koneksi ke MongoDB (periksa variabel lingkungan yang digunakan untuk URL MongoDB)
	database.MongoConnect()

	// Setup semua route
	mux := routes.SetupRoutes()

	// Ambil PORT dari environment, default ke 8080 jika tidak ditemukan
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback ke 8080 jika PORT tidak ditemukan
	}

	// Jalankan server di port yang sesuai
	fmt.Println("Server is running on http://localhost:" + port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
