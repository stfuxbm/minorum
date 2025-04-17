package routes

import (
	"net/http"

	ordosHandlers "github.com/stfuxbm/quote-saints/internal/handlers/ordos"
	quotesHandlers "github.com/stfuxbm/quote-saints/internal/handlers/quotes"
	"github.com/stfuxbm/quote-saints/internal/middleware"
)

func SetupRoutes() http.Handler {
	// Buat router bawaan Go
	mux := http.NewServeMux()

	// routes untuk quote
	mux.HandleFunc("/api/v1/quotes", quotesHandlers.AddQuote)
	mux.HandleFunc("/api/v1/random-quotes", quotesHandlers.GetRandomQuotes)
	mux.HandleFunc("/api/v1/quotes/search", quotesHandlers.GetQuotesBySaintName)
	mux.HandleFunc("/api/v1/quotes/category", quotesHandlers.GetQuotesByCategory)

	// routes untuk ordo
	mux.HandleFunc("/api/v1/ordos", ordosHandlers.AddOrdo)
	mux.HandleFunc("/api/v1/search", ordosHandlers.GetOrdoByOrderNameOrNickname)

	// middleware untuk logging dan CORS
	return middleware.CORS(
		middleware.Logger(mux),
	)
}
