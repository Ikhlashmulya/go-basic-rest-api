package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ikhlashmulya/go-basic-rest-api/database"
	"github.com/ikhlashmulya/go-basic-rest-api/model"
	"github.com/julienschmidt/httprouter"
)

func GetStudents(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	students := database.FindAll()

	bytes, err := json.Marshal(&students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JsonResponse(w, bytes)
}

func SaveStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	student := new(model.Student)

	err = json.Unmarshal(payload, student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.Save(*student)
	w.Header().Add("Location", r.URL.Path+"/"+student.Id)
}

func GetStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	student, ok := database.FindById(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	bytes, err := json.Marshal(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JsonResponse(w, bytes)
}

func JsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("content-type", "application/json")
	w.Write(bytes)
}
