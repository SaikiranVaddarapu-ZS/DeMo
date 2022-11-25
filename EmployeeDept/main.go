package main

import (
	"EmployeeDept/Emp"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateDB() (*sql.DB, error) {
	DB, er := sql.Open("mysql", "root:Saikiran@18@tcp(127.0.0.1:3306)/employeedata")
	if er != nil {
		return nil, er
	}
	return DB, nil
}

func main() {
	var er error
	Emp.DB, er = sql.Open("mysql", "root:Saikiran@18@tcp(127.0.0.1:3306)/employeedata")
	//DB,er := CreateDB()
	if er != nil {
		log.Println(er)
		return
	}
	defer Emp.DB.Close()
	if er != nil {
		log.Println(er)
		return
	}
	router := mux.NewRouter()
	router.HandleFunc("/employee", Emp.GetEmployeeHandler).Methods("GET")
	router.HandleFunc("/allemployees", Emp.GetAllEmployeeHandler).Methods("GET")
	router.HandleFunc("/addemployee", Emp.PostEmployeeHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
