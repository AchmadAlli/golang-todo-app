package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type RestApi struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

func Response(writer http.ResponseWriter, code int, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)

	body := RestApi{
		Status: "OK",
		Data:   data,
		Error:  nil,
	}

	json.NewEncoder(writer).Encode(body)
}

func ErrorResponse(writer http.ResponseWriter, code int, message string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)

	body := RestApi{
		Status: "Error",
		Data:   nil,
		Error:  message,
	}

	json.NewEncoder(writer).Encode(body)
}

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
