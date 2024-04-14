package response

import (
	"github.com/goccy/go-json"
	"net/http"
)

// Response schema
type Response struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return
		}
	}
}
