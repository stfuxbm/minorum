package handlers

import (
	"net/http"

	helpers "github.com/stfuxbm/minorum/internal/helpers/saints"
	models "github.com/stfuxbm/minorum/internal/models/response"
	"github.com/stfuxbm/minorum/internal/utils"
)

// handler post

func AddSaint(w http.ResponseWriter, r *http.Request) {

	// validasi only post method
	if r.Method != http.MethodPost {
		utils.ErrorResponse(
			w,
			models.MsgMethodNotAllowed,
			http.StatusMethodNotAllowed,
		)
	}

	// decode and validasi input
	saint, err := helpers.DecodeSaint(r) // ganti method dan variabel
	if err != nil {
		utils.ErrorResponse(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}
	// check filed are required
	if err := helpers.ValidateSaint(saint); // ganti method dan variabel
	err != nil {
		utils.ErrorResponse(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}
	// check duplicate
	if err := helpers.CheckDuplicateSaint(r, saint); // ganti method dan variabel
	err != nil {
		utils.ErrorResponse(
			w,
			err.Error(),
			http.StatusConflict,
		)
		return
	}

	// save to database
	if err := helpers.SaveSaint(r, &saint); // ganti method dan variabel
	err != nil {
		utils.ErrorResponse(
			w,
			models.MsgInternalServerError,
			http.StatusInternalServerError,
		)
	}
	utils.SuccessResponse(
		w,
		nil,
		"Saint successfully added",
		http.StatusCreated,
	)
}
