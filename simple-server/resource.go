package main

import (
	"encoding/json"
	"net/http"
)

func process(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		var contact = Contact{
			ID:    1,
			Name:  "Thom",
			City:  "Lima",
			Phone: "chismos@ eres >:v",
		}
		// res.Header().Set("Content-type", "text/plain")
		// res.Header().Set("Content-type", "application/json")
		// res.Write([]byte(contact.toString()))
		js, err := json.Marshal(contact)
		res.Header().Set("Content-Type", "application/json")
		if err != nil {
			return
		}
		res.Write(js)
	}
}
