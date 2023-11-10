package utils

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Data any `json:"data"`
}

func WriteJSON(w http.ResponseWriter, data any) {
	res := JSONResponse{Data: data}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
