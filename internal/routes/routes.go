package routes

import (
	"net/http"

	ordosHandlers "github.com/stfuxbm/minorum/internal/handlers/ordos"
	quotesHandlers "github.com/stfuxbm/minorum/internal/handlers/quotes"
	saintHandlers "github.com/stfuxbm/minorum/internal/handlers/saints"
	"github.com/stfuxbm/minorum/internal/middleware"
)

func SetupRoutes() http.Handler {
	// Buat router bawaan Go
	mux := http.NewServeMux()

	// routes untuk quote
	mux.HandleFunc("/api/v1/quotes", quotesHandlers.AddQuote)                              // Menambahkan quote
	mux.HandleFunc("/api/v1/random-quotes", quotesHandlers.GetRandomQuotes)                // Mendapatkan quote acak
	mux.HandleFunc("/api/v1/quotes/search", quotesHandlers.GetQuotesBySaintNameOrCategory) // Mencari quote berdasarkan nama atau kategori

	// routes untuk ordo
	mux.HandleFunc("/api/v1/ordos", ordosHandlers.AddOrdo)                             // Menambahkan ordo
	mux.HandleFunc("/api/v1/ordos/search", ordosHandlers.GetOrdoByOrderNameOrNickname) // Mencari ordo berdasarkan nama atau nickname

	// routes untuk santo
	mux.HandleFunc("/api/v1/saints", saintHandlers.AddSaint)                             // Menambahkan santo
	mux.HandleFunc("/api/v1/saints/search", saintHandlers.GetSaintByOrderNameOrNickname) // Mencari santo berdasarkan nama ordo atau nickname

	return middleware.CORS( // cors
		middleware.Logger( // logger
			middleware.RateLimit(mux), // Apply rate limiting middleware
		),
	)
}
