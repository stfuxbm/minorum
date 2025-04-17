package handlers

import (
	"net/http"
	"strings"

	helpers "github.com/stfuxbm/quote-saints/internal/helpers/ordos"
	models "github.com/stfuxbm/quote-saints/internal/models/response"

	"github.com/stfuxbm/quote-saints/internal/utils"
)

// GetOrdoByOrderNameOrNickname menangani permintaan GET untuk mendapatkan ordo berdasarkan order_name atau nickname.
func GetOrdoByOrderNameOrNickname(w http.ResponseWriter, r *http.Request) {
	// Hanya izinkan method GET
	if r.Method != http.MethodGet {
		utils.ErrorResponse(
			w,
			models.MsgMethodNotAllowed,
			http.StatusMethodNotAllowed,
		)
		return
	}

	// Ambil query parameter 'order' dari URL
	order := r.URL.Query().Get("order")
	if order == "" {
		utils.ErrorResponse(
			w,
			"Order parameter is required",
			http.StatusBadRequest,
		)
		return
	}

	// Normalisasi order_name dan nickname (agar pencarian lebih fleksibel)
	order = strings.ToLower(order)

	// Ambil ordo berdasarkan order_name atau nickname dari database
	ordo, err := helpers.GetOrdoByOrderName(r.Context(), order)
	if err != nil {
		utils.ErrorResponse(
			w,
			"Failed to retrieve order",
			http.StatusInternalServerError,
		)
		return
	}

	// Jika tidak ada ordo untuk order_name atau nickname yang diberikan, kembalikan response error
	if ordo == nil {
		utils.ErrorResponse(
			w,
			"No order found for the given order_name or nickname",
			http.StatusNotFound,
		)
		return
	}

	// Kirim response dengan ordo yang ditemukan
	utils.SuccessResponse(
		w,
		ordo,
		"Order successfully retrieved",
		http.StatusOK,
	)
}
