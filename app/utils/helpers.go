package utils

import (
	"encoding/json"
	"net/http"
)

func Response(writer http.ResponseWriter, body interface{}) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(body)
}

func ErrorResponse(writer http.ResponseWriter, code int, message string) {
	writer.WriteHeader(code)
	writer.Header().Set("Content-Type", "application/json")
	body := map[string]string{
		"error": message,
	}

	json.NewEncoder(writer).Encode(body)
}
