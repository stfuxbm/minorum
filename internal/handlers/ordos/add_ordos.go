package handlers

import (
	"net/http"

	helpers "github.com/stfuxbm/minorum/internal/helpers/ordos"
	models "github.com/stfuxbm/minorum/internal/models/response"
	"github.com/stfuxbm/minorum/internal/utils"
)

// AddOrdo menangani permintaan POST untuk menambahkan ordo baru ke koleksi "ordos".
func AddOrdo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, models.MsgMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	// Decode dan validasi request
	ordo, err := helpers.DecodeOrdo(r)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := helpers.ValidateOrdo(ordo); err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Cek duplikasi berdasarkan nama
	if err := helpers.CheckDuplicateOrdo(r, ordo); err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusConflict)
		return
	}

	// Simpan data ke database
	if err := helpers.SaveOrdo(r, &ordo); err != nil {
		utils.ErrorResponse(w, models.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	// Sukses
	utils.SuccessResponse(w, nil, "Ordo successfully added", http.StatusCreated)
}
