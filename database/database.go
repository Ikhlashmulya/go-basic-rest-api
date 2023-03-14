package database

import "github.com/ikhlashmulya/go-basic-rest-api/model"

var database = make(map[string]model.Student)

func FindAll() []model.Student {
	database["01"] = model.Student{
		Id:    "01",
		Name:  "Ikhlash Mulyanurahman",
		Email: "nurahmanmulya@gmail.com",
	}

	items := make([]model.Student, 0)
	for _, val := range database {
		items = append(items, val)
	}

	return items
}

func Save(key string, data model.Student) {
	database[key] = data
}

func FindById(id string) (model.Student, bool) {
	student, ok := database[id]

	return student, ok
}

func Remove(id string) {
	delete(database, id)
}
