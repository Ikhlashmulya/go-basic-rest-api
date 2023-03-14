package main

import (
	"log"
	"net/http"

	"github.com/ikhlashmulya/go-basic-rest-api/handler"
	"github.com/julienschmidt/httprouter"
)

func main() {

	//router
	router := httprouter.New()
	router.GET("/students", handler.GetStudents)

	//server
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server Running on 8080")
	err := server.ListenAndServe()
	log.Fatal(err)
}
