package routes

import (
	"net/http"

	"github.com/stfuxbm/quote-saints/internal/handlers"
	"github.com/stfuxbm/quote-saints/internal/middleware"
)

func SetupRoutes() http.Handler {
	// Buat router bawaan Go
	mux := http.NewServeMux()

	//  route untuk menambah quote baru
	mux.HandleFunc("/api/v1/quotes", handlers.AddQuote)

	//  route untuk mendapatkan quote acak
	mux.HandleFunc("/api/v1/random-quotes", handlers.GetRandomQuotes)

	//  route untuk mendapatkan quote berdasarkan nama santo
	mux.HandleFunc("/api/v1/quotes/search", handlers.GetQuotesBySaintName)

	//  route untuk mendapatkan quote berdasarkan kategori
	mux.HandleFunc("/api/v1/quotes/category", handlers.GetQuotesByCategory)

	//  middleware untuk logging dan CORS
	return middleware.CORS(
		middleware.Logger(mux),
	)
}
