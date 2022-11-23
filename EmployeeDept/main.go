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
	router.HandleFunc("/emp", Emp.GetEmployeeHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
