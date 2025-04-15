package handlers

import (
	"net/http"

	"github.com/stfuxbm/quote-saints/internal/helpers"
	"github.com/stfuxbm/quote-saints/internal/models"
	"github.com/stfuxbm/quote-saints/internal/utils"
)

// AddQuote menangani permintaan POST untuk menambahkan quote baru ke koleksi "quotes".
func AddQuote(w http.ResponseWriter, r *http.Request) {
	// Hanya izinkan method POST
	if r.Method != http.MethodPost {
		utils.ErrorResponse(
			w,
			models.MsgMethodNotAllowed,
			http.StatusMethodNotAllowed,
		)
		return
	}

	// Decode body menjadi struct Quote
	quote, err := helpers.DecodeQuote(r)
	if err != nil {
		utils.ErrorResponse(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	// Validasi isi data quote
	if err := helpers.ValidateQuote(quote); err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Cek apakah quote yang sama sudah ada
	if err := helpers.CheckDuplicateQuote(r, quote); err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusConflict)
		return
	}

	// Simpan quote ke MongoDB
	if err := helpers.SaveQuote(r, &quote); err != nil {
		utils.ErrorResponse(w, models.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	// Kirim response sukses
	utils.SuccessResponse(
		w,
		nil,
		"Quote successfully added",
		http.StatusCreated,
	)
}
