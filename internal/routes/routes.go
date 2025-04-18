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
	mux.HandleFunc("/api/v1/quotes", quotesHandlers.AddQuote)
	mux.HandleFunc("/api/v1/random-quotes", quotesHandlers.GetRandomQuotes)
	mux.HandleFunc("/api/v1/quotes/search", quotesHandlers.GetQuotesBySaintNameOrCategory)

	// routes untuk ordo
	mux.HandleFunc("/api/v1/ordos", ordosHandlers.AddOrdo)
	mux.HandleFunc("/api/v1/ordos/search", ordosHandlers.GetOrdoByOrderNameOrNickname)

	// routes for apostle
	mux.HandleFunc("/api/v1/saints", saintHandlers.AddSaint)
	mux.HandleFunc("/api/v1/saints/search", saintHandlers.GetSaintByOrderNameOrNickname)
	// middleware untuk logging dan CORS
	return middleware.CORS(
		middleware.Logger(mux),
	)
}
