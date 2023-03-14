package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8000",
	}

	err := server.ListenAndServe()
	log.Fatal(err)
}
