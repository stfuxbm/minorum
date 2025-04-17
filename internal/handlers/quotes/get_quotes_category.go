package handlers

import (
	"net/http"

	helpers "github.com/stfuxbm/quote-saints/internal/helpers/quotes"
	models "github.com/stfuxbm/quote-saints/internal/models/response"

	"github.com/stfuxbm/quote-saints/internal/utils"
)

// GetQuotesByCategory menangani permintaan GET untuk mendapatkan quote berdasarkan kategori.
func GetQuotesByCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorResponse(w, models.MsgMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	// Ambil semua query param 'category'
	categories := r.URL.Query()["category"]
	if len(categories) == 0 {
		utils.ErrorResponse(w, "At least one category parameter is required", http.StatusBadRequest)
		return
	}

	// Ambil quotes berdasarkan beberapa kategori
	quotes, err := helpers.GetQuotesByCategories(r.Context(), categories)
	if err != nil {
		utils.ErrorResponse(w, "Failed to retrieve quotes", http.StatusInternalServerError)
		return
	}

	if len(quotes) == 0 {
		utils.ErrorResponse(w, "No quotes found for the given categories", http.StatusNotFound)
		return
	}

	utils.SuccessResponse(w, quotes, "Quotes retrieved for the selected categories", http.StatusOK)
}
