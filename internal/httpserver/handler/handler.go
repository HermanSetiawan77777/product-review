package handler

import (
	"encoding/json"
	"net/http"

	review "github.com/HermanSetiawan77777/JakMall/internal/service/review"
)

func GetSummary(w http.ResponseWriter, r *http.Request) {

	summary := review.CalculatedSummary()
	responseWithData(w, http.StatusOK, summary)
}

func responseWithData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
