package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/owais-webonise/restapi/models"
)

func Getstudents(w http.ResponseWriter, r *http.Request) {
	students := models.Students
	json.NewEncoder(w).Encode(students)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range models.Students {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Student{})
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	students := models.Students
	var Student models.Student
	_ = json.NewDecoder(r.Body).Decode(&Student)
	students = append(students, Student)
	json.NewEncoder(w).Encode(students)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	students := models.Students
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(students)
	}
}
