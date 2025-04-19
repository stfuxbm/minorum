package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// getIPAddress mengambil alamat IP pengunjung dari header HTTP
func getIPAddressAccess(r *http.Request) string {
	// Memeriksa beberapa header yang dapat berisi alamat IP
	// X-Forwarded-For digunakan jika aplikasi berjalan di belakang proxy atau load balancer
	ips := strings.Split(r.Header.Get("X-Forwarded-For"), ",")
	if len(ips) > 0 {
		return strings.TrimSpace(ips[0]) // Mengambil IP pertama
	}

	// Jika X-Forwarded-For tidak ada, ambil IP langsung dari koneksi
	return r.RemoteAddr
}

// Logger adalah middleware untuk mencatat semua request yang masuk, termasuk IP, waktu, metode, dan path.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil IP pengunjung
		ip := getIPAddressAccess(r)

		// Ambil waktu permintaan
		requestTime := time.Now().Format("Jan 02 03:04:05 PM")

		// Ambil metode HTTP (GET, POST, dll)
		method := r.Method

		// Ambil path dari URL
		path := r.URL.Path

		// Ambil User-Agent
		userAgent := r.Header.Get("User-Agent")

		// Mencatat log dengan format yang diinginkan
		log.Printf("%s | Method: %s | Path: %s | IP: %s | User-Agent: %s", requestTime, method, path, ip, userAgent)

		// Lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}
