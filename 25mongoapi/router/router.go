package router

import (
	"github.com/fahimishraq99/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/students", controller.GetAllStudents).Methods("GET")
	router.HandleFunc("/api/student", controller.CreateStudent).Methods("POST")
	router.HandleFunc("/api/student/{id}", controller.MarkAsStudied).Methods("PUT")
	router.HandleFunc("/api/student/{id}", controller.DeleteAStudent).Methods("DELETE")
	router.HandleFunc("/api/deleteallstudent", controller.DeleteAllStudents).Methods("DELETE")

	return router
}
