package routes

import (
	"net/http"

	"github.com/stfuxbm/quote-saints/internal/handlers"
	"github.com/stfuxbm/quote-saints/internal/middleware"
)

func SetupRoutes() http.Handler {
	// Buat router bawaan Go
	mux := http.NewServeMux()

	// Daftarkan route untuk menambah quote baru
	mux.HandleFunc("/api/v1/quotes", handlers.AddQuote)

	// Daftarkan route untuk mendapatkan quote acak
	mux.HandleFunc("/api/v1/random-quotes", handlers.GetRandomQuotes)

	// Daftarkan route untuk mendapatkan quote berdasarkan nama santo
	mux.HandleFunc("/api/v1/quotes/search", handlers.GetQuotesBySaintName)

	// Tambahkan middleware untuk logging dan CORS
	// Logger akan menangani request dan mencatat log

	return middleware.CORS(
		middleware.Logger(mux), // Logger middleware
	)
}
