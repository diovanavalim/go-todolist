package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func prepare(w http.ResponseWriter, statusCode int) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return w
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	responseWriter := prepare(w, statusCode)

	if data != nil {
		if err := json.NewEncoder(responseWriter).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}{
		Message: err.Error(),
		Status:  statusCode,
	})
}
