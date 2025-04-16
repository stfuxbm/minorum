package handlers

import (
	"net/http"
	"strings"

	"github.com/stfuxbm/quote-saints/internal/helpers"
	"github.com/stfuxbm/quote-saints/internal/models"
	"github.com/stfuxbm/quote-saints/internal/utils"
)

// GetQuotesBySaintName menangani permintaan GET untuk mendapatkan quote berdasarkan nama santo.
func GetQuotesBySaintName(w http.ResponseWriter, r *http.Request) {
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
	name := r.URL.Query().Get("name")
	if name == "" {
		utils.ErrorResponse(
			w,
			"Name parameter is required",
			http.StatusBadRequest,
		)
		return
	}

	// Normalisasi nama (agar pencarian lebih fleksibel)
	name = strings.ToLower(name)

	// Ambil quotes berdasarkan nama santo dari database
	quotes, err := helpers.GetQuotesBySaintName(r.Context(), name)
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
			"No quotes found for the given saint name",
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
