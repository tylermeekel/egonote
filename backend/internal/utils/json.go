package utils

import (
	"encoding/json"
	"net/http"
)

type JSONResponseWriter struct {
	Writer http.ResponseWriter `json:"-"`
	Data   map[string]any      `json:"data,omitempty"`
	Errors map[string][]string `json:"errors,omitempty"`
}

func (jw *JSONResponseWriter) AddData(key string, value any) {
	if len(jw.Errors) == 0 {
		jw.Data[key] = value
	}
}

func (jw *JSONResponseWriter) AddError(key, value string) {
	jw.Errors[key] = append(jw.Errors[key], value)
}

func (jw *JSONResponseWriter) AddInternalError() {
	if _, ok := jw.Errors["internal"]; !ok {
		jw.Errors["internal"] = []string{"internal server error"}
	}
}

func (jw *JSONResponseWriter) WriteJSON() {
	jw.Writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(jw.Writer).Encode(jw)
}

func NewJSONResponseWriter(w http.ResponseWriter) JSONResponseWriter {
	return JSONResponseWriter{
		Writer: w,
		Data:   make(map[string]any),
		Errors: make(map[string][]string),
	}
}
