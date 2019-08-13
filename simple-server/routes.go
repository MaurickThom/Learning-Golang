package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func startRoutes() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
