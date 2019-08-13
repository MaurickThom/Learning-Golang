package main

import "net/http"

func process(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		var contact = Contact{
			ID:    1,
			Name:  "Thom",
			City:  "Lima",
			Phone: "chismos@ eres >:v",
		}
		res.Header().Set("Content-type", "text/plain")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(contact.toString()))
	}
}
