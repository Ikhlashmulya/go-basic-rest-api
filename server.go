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
	router.GET("/api/students", handler.GetStudents)
	router.POST("/api/students", handler.SaveStudent)

	router.GET("/api/students/:id", handler.GetStudent)
	router.PUT("/api/students/:id", handler.UpdateStudent)
	router.DELETE("/api/students/:id", handler.DeleteStudent)

	//server
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server Running on 8080")
	err := server.ListenAndServe()
	log.Fatal(err)
}
