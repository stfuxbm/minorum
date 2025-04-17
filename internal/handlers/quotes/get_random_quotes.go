package handlers

import (
	"math/rand"
	"net/http"
	"time"

	helpers "github.com/stfuxbm/quote-saints/internal/helpers/quotes"
	models "github.com/stfuxbm/quote-saints/internal/models/response"

	"github.com/stfuxbm/quote-saints/internal/utils"
)

// GetRandomQuote menangani permintaan GET untuk mendapatkan quote acak.
func GetRandomQuotes(w http.ResponseWriter, r *http.Request) {
	// Hanya izinkan method GET
	if r.Method != http.MethodGet {
		utils.ErrorResponse(
			w,
			models.MsgMethodNotAllowed,
			http.StatusMethodNotAllowed,
		)
		return
	}

	// Ambil semua quote dari database
	quotes, err := helpers.GetRandomQuotes(r.Context())
	if err != nil {
		utils.ErrorResponse(
			w,
			"Failed to retrieve quotes",
			http.StatusInternalServerError,
		)
		return
	}

	// Jika tidak ada quotes, kembalikan response error
	if len(quotes) == 0 {
		utils.ErrorResponse(
			w,
			"No quotes available",
			http.StatusNotFound,
		)
		return
	}

	// Pilih satu quote secara acak
	rand.Seed(time.Now().UnixNano())      // Inisialisasi seed untuk random
	randomIndex := rand.Intn(len(quotes)) // Pilih index acak dalam range quotes
	randomQuote := quotes[randomIndex]    // Ambil quote acak berdasarkan index

	// Kirim response dengan quote acak
	utils.SuccessResponse(
		w,
		randomQuote,
		"Random quote successfully retrieved",
		http.StatusOK,
	)
}
