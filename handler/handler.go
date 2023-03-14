package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ikhlashmulya/go-basic-rest-api/database"
	"github.com/julienschmidt/httprouter"
)

func GetStudents(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	students := database.FindAll()

	bytes, err := json.Marshal(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JsonResponse(w, bytes, http.StatusOK)
}

func JsonResponse(w http.ResponseWriter, bytes []byte, code int) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(bytes)
}
