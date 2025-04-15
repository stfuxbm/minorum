package middleware

import (
	"log"
	"net/http"
)

// Logger adalah middleware untuk mencatat semua request yang masuk
// Middleware ini berguna untuk mencatat informasi setiap permintaan HTTP yang masuk ke server
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mencatat method HTTP dan URL path dari request yang diterima
		// Contoh log: Request: GET /api/v1/new-testament/chapters
		log.Printf("Request: %s %s", r.Method, r.URL.Path)

		// Lanjutkan ke handler berikutnya setelah mencatat log
		// Middleware ini tidak menghalangi request untuk diteruskan, hanya mencatatnya
		next.ServeHTTP(w, r)
	})
}
