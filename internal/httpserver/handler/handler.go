package handler

import (
	"encoding/json"
	"net/http"

	review "github.com/HermanSetiawan77777/JakMall/internal/service/review"
	"github.com/gorilla/mux"
)

func GetSummary(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	summary := review.CalculatedSummary(name)
	responseWithData(w, http.StatusOK, summary)
}

func responseWithData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
