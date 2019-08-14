package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()
	router.Handle("/", http.FileServer(http.Dir("./public")))
	var server = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
