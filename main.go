package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ProFL/viability-challenge/viability"
)

func main() {
	r := viability.GetRouter()

	server := &http.Server{
		Addr:         "0.0.0.0:3000",
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	log.Panic(server.ListenAndServe())
}
