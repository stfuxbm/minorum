package routes

import (
	"net/http"

	"github.com/stfuxbm/quote-saints/internal/handlers"
	"github.com/stfuxbm/quote-saints/internal/middleware"
)

func SetupRoutes() http.Handler {
	// Buat router bawaan Go
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/quotes", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddQuote(w, r)
	})

	// Tambahkan middleware untuk logging dan CORS
	return middleware.CORS(
		middleware.Logger(mux),
	)
}
