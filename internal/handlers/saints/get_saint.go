package handlers

import (
	"net/http"
	"strings"

	helpers "github.com/stfuxbm/minorum/internal/helpers/saints" // ganti path helpers sesuai
	models "github.com/stfuxbm/minorum/internal/models/response"
	"github.com/stfuxbm/minorum/internal/utils"
)

// GetSaintByOrderNameOrNickname menangani permintaan GET untuk mendapatkan saint berdasarkan name atau nickname.
func GetSaintByOrderNameOrNickname(w http.ResponseWriter, r *http.Request) {
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

	// Normalisasi nama dan nickname untuk pencarian fleksibel
	order = strings.ToLower(order)

	// Ambil saint berdasarkan name atau nickname dari database
	saint, err := helpers.GetSaintByNameOrNickname(r.Context(), order) // ganti method dan variabel
	if err != nil {
		utils.ErrorResponse(
			w,
			"Failed to retrieve saint",
			http.StatusInternalServerError,
		)
		return
	}

	// Jika tidak ditemukan
	if saint == nil {
		utils.ErrorResponse(
			w,
			"No saint found for the given name or nickname",
			http.StatusNotFound,
		)
		return
	}

	// Kirim response sukses
	utils.SuccessResponse(
		w,
		saint,
		"Saint successfully retrieved",
		http.StatusOK,
	)
}
