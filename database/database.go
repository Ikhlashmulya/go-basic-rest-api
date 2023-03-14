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

func Save(data model.Student) {
	database[data.Id] = data
}

func FindById(id string) (model.Student, bool) {
	student, ok := database[id]

	return student, ok
}
