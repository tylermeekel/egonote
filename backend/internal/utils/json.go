package utils

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Data any `json:"data"`
}

type JSONErrorResponse struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, data any) {
	res := JSONResponse{Data: data}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func WriteJSONError(w http.ResponseWriter, err string) {
	res := JSONErrorResponse{Error: err}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func WriteInternalServerError(w http.ResponseWriter) {
	res := JSONErrorResponse{Error: "internal server error"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(res)
}
