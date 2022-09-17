package main

import (
	"fmt"
	"net/http"

	"github.com/HermanSetiawan77777/JakMall/internal/httpserver"
)

func main() {
	appserver := &http.Server{
		Addr:    "localhost:8080",
		Handler: httpserver.HandleRoutes(),
	}
	fmt.Println("Server Port 8080 is run !")
	appserver.ListenAndServe()

}
