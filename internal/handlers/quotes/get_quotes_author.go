package handlers

import (
	"net/http"
	"strings"

	helpers "github.com/stfuxbm/minorum/internal/helpers/quotes"
	models "github.com/stfuxbm/minorum/internal/models/response"
	"github.com/stfuxbm/minorum/internal/utils"
)

// GetQuotesBySaintName menangani permintaan GET untuk mendapatkan quote berdasarkan nama santo.
func GetQuotesBySaintNameOrCategory(w http.ResponseWriter, r *http.Request) {
	// Hanya izinkan method GET
	if r.Method != http.MethodGet {
		utils.ErrorResponse(
			w,
			models.MsgMethodNotAllowed,
			http.StatusMethodNotAllowed,
		)
		return
	}

	// Ambil query parameter 'name' dari URL
	order := r.URL.Query().Get("order")
	if order == "" {
		utils.ErrorResponse(
			w,
			"Order parameter is required",
			http.StatusBadRequest,
		)
		return
	}

	// Normalisasi nama (agar pencarian lebih fleksibel)
	order = strings.ToLower(order)

	// Ambil quotes berdasarkan nama santo dari database
	quotes, err := helpers.GetQuotesBySaintName(r.Context(), order)
	if err != nil {
		utils.ErrorResponse(
			w,
			"Failed to retrieve quotes",
			http.StatusInternalServerError,
		)
		return
	}

	// Jika tidak ada quotes untuk nama santo yang diberikan, kembalikan response error
	if len(quotes) == 0 {
		utils.ErrorResponse(
			w,
			"No quotes found for the given saint name or category",
			http.StatusNotFound,
		)
		return
	}

	// Kirim response dengan quotes yang ditemukan
	utils.SuccessResponse(
		w,
		quotes,
		"Quotes successfully retrieved for the saint",
		http.StatusOK,
	)
}
