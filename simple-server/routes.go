package main

import (
	"fmt"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello world")
}

func startRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/contacts", process)
	http.ListenAndServe(":8080", nil)
}
