package httpserver

import (
	"net/http"

	"github.com/HermanSetiawan77777/JakMall/internal/httpserver/handler"
	"github.com/gorilla/mux"
)

func HandleRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/summary/", handler.GetSummary).Methods(http.MethodGet)
	r.HandleFunc("/summary/{name}", handler.GetSummary).Methods(http.MethodGet)
	return r
}
