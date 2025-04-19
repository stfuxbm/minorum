package middleware

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	models "github.com/stfuxbm/minorum/internal/models/response"
	"golang.org/x/time/rate"
)

// Visitor menyimpan data rate limiter dan timestamp terakhir
type Visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	visitors = make(map[string]*Visitor)
	mu       sync.Mutex
	// Konfigurasi rate limiter
	rateLimit       = 1.0 // 1 request per detik
	burstSize       = 2   // Maksimal 2 request pada waktu yang sama
	cleanupInterval = 5 * time.Minute
)

// Jalankan pembersihan visitor yang sudah tidak aktif
func init() {
	go cleanupVisitors()
}

// Membersihkan visitor yang tidak aktif dalam jangka waktu tertentu
func cleanupVisitors() {
	for {
		time.Sleep(cleanupInterval)

		mu.Lock()
		for ip, visitor := range visitors {
			if time.Since(visitor.lastSeen) > 10*time.Minute {
				delete(visitors, ip)
				log.Printf("Cleaned up rate limiter for inactive IP: %s", ip)
			}
		}
		mu.Unlock()
	}
}

// getVisitor mengembalikan rate limiter untuk IP tertentu.
// Jika belum ada, dibuatkan limiter baru.
func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	visitor, exists := visitors[ip]
	if !exists {
		log.Printf("Creating new rate limiter for IP: %s (rate: %.1f/s, burst: %d)", ip, rateLimit, burstSize)
		limiter := rate.NewLimiter(rate.Limit(rateLimit), burstSize)
		visitor = &Visitor{
			limiter:  limiter,
			lastSeen: time.Now(),
		}
		visitors[ip] = visitor
	} else {
		// Update timestamp terakhir dilihat
		visitor.lastSeen = time.Now()
	}

	return visitor.limiter
}

// getIPAddress mencoba mendapatkan IP asli dari request.
// Pertama mencoba header X-Forwarded-For, jika tidak ada, fallback ke RemoteAddr.
func getIPAddress(r *http.Request) string {
	// Cek apakah ada header X-Forwarded-For
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		// Jika ada beberapa IP, ambil yang pertama
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// Jika tidak ada X-Forwarded-For, fallback ke RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // fallback jika terjadi error
	}
	return ip
}

// RateLimit adalah middleware untuk membatasi jumlah permintaan berdasarkan IP per detik.
// Jika melebihi batas, akan mengembalikan status 429 dengan respons JSON.
func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIPAddress(r)
		path := r.URL.Path

		log.Printf("[RateLimit] Request from IP: %s, Path: %s", ip, path)

		limiter := getVisitor(ip)

		// Ambil status token dari limiter
		tokensRemaining := limiter.Tokens()

		// Cek apakah permintaan diizinkan
		allowed := limiter.Allow()
		log.Printf("[RateLimit] IP: %s, Path: %s, Tokens Remaining: %.2f, Request Allowed: %v",
			ip, path, tokensRemaining, allowed)

		if !allowed {
			log.Printf("[RateLimit] BLOCKED - Rate limit exceeded for IP: %s, Path: %s", ip, path)

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Retry-After", "1") // Sarankan client untuk mencoba lagi setelah 1 detik
			w.WriteHeader(http.StatusTooManyRequests)

			// Kirim response JSON
			resp := models.Response{
				Success: false,
				Message: models.MsgTooManyRequests,
				Code:    http.StatusTooManyRequests,
				Data:    nil,
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Log bahwa permintaan diizinkan
		log.Printf("[RateLimit] ALLOWED - Request permitted for IP: %s, Path: %s", ip, path)

		// Lanjutkan ke handler berikutnya jika tidak melebihi limit
		next.ServeHTTP(w, r)
	})
}
