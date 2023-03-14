package database

import (
	"testing"

	"github.com/ikhlashmulya/go-basic-rest-api/model"
)

var student = model.Student{
	Id:    "01",
	Name:  "Ikhlash Mulyanurahman",
	Email: "imulya664@gmail.com",
}

func TestSaveStudent(t *testing.T) {
	Save(student.Id, student)

	_, ok := FindById(student.Id)

	if !ok {
		t.Error("data is not saved")
	}
}

func TestGetStudent(t *testing.T) {
	_, ok := FindById("01")

	if !ok {
		t.Error("data not found")
	}
}

func TestDeleteStudent(t *testing.T) {
	Remove(student.Id)

	_, ok := FindById(student.Id)

	if ok {
		t.Error("data is not deleted")
	}
}
