package main

import (
	"EmployeeDept/Emp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employee", Emp.GetEmployeeHandler).Methods("GET")
	router.HandleFunc("/allemployees", Emp.GetAllEmployeeHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
