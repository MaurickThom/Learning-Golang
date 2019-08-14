package main

import "net/http"

func main() {
	var mux = http.NewServeMux()                 // mutiplexor o roteador
	var fs = http.FileServer(http.Dir("public")) //public
	mux.Handle("/", fs)
	http.ListenAndServe(":8080", mux)
}
