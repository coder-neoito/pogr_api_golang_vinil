package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vinilunni/pogr_api_golang_vinil/internal/common/types"
)

func RespondWithJSON(w http.ResponseWriter, status int, dataPayload types.ResponseTemplate) {
	resp, err := json.Marshal(dataPayload)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Server Error!")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(resp)
}

func RespondWithError(w http.ResponseWriter, status int, message string) {
	var resp types.ResponseTemplate
	resp.Success = false
	resp.Message = message
	errResp, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(errResp)
}

func RespondWithNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
