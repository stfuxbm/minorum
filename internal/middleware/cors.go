package middleware

import (
	"net/http"
)

// CORS adalah middleware untuk menambahkan header CORS
// CORS (Cross-Origin Resource Sharing) memungkinkan resource dari server yang berbeda untuk diakses oleh aplikasi frontend.
// Biasanya digunakan untuk mengizinkan akses dari domain yang berbeda.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menambahkan header "Access-Control-Allow-Origin" untuk mengizinkan semua domain (dengan "*")
		// Biasanya kita bisa spesifik menyebutkan domain yang diizinkan, tetapi "*" berarti semua domain diizinkan.
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Menambahkan header "Access-Control-Allow-Methods" untuk mengizinkan metode HTTP tertentu
		// Dalam hal ini, hanya metode GET, POST, PUT, DELETE yang diizinkan untuk akses dari domain lain
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

		// Menambahkan header "Access-Control-Allow-Headers" untuk mengizinkan header tertentu dalam request
		// Dalam hal ini, hanya Content-Type yang diizinkan
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Jika request yang diterima adalah OPTIONS (preflight request), kita kirimkan response dengan status OK (200)
		// OPTIONS request digunakan oleh browser untuk mengecek apakah metode yang diminta diizinkan oleh server
		// Biasanya digunakan sebelum melakukan permintaan dengan metode seperti POST atau PUT.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK) // Response OK untuk OPTIONS request
			return
		}

		// Jika bukan OPTIONS request, lanjutkan ke handler berikutnya dalam chain middleware
		// Ini memastikan request diteruskan setelah header CORS ditambahkan
		next.ServeHTTP(w, r)
	})
}
