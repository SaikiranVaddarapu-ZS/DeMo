package Emp

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var mock sqlmock.Sqlmock
var err error

func TestGetEmployeeHandler(t *testing.T) {
	DB, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println(err)
		return
	}
	defer DB.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "phone", "deptid", "name"}).
		AddRow("21", "Sarath", "5467890321", "121", "Kroger")

	mock.ExpectQuery("select e.id,e.name,e.phone,e.deptid,d.name from Employee e join Department d on e.deptid = d.id where e.id = ?").WithArgs("21").WillReturnRows(rows)
	req, err := http.NewRequest("GET", "http://localhost:8080/employee?id=21", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEmployeeHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Id":"21","Name":"Sarath","Phone":"5467890321","Dept":{"Id":"121","Name":"Kroger"}}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got (%v) want (%v)",
			rr.Body.String(), expected)
	}
}

func TestGetAllEmployeeHandler(t *testing.T) {
	DB, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println(err)
		return
	}
	defer DB.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "phone", "deptid", "name"}).
		AddRow("21", "Sarath", "5467890321", "121", "Kroger")

	mock.ExpectQuery("select e.id,e.name,e.phone,e.deptid,d.name from Employee e join Department d on e.deptid = d.id").WillReturnRows(rows)
	req, err := http.NewRequest("GET", "http://localhost:8080/allemployees", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllEmployeeHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPostEmployeeHandler(t *testing.T) {
	DB, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println(err)
		return
	}
	defer DB.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO employee").WithArgs(sqlmock.AnyArg(), "Hena", "121", "8765490123").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	em := Emp{"", "Hena", "8765490123", Dept{"121", ""}}
	empbody, _ := json.Marshal(em)
	req, err := http.NewRequest("POST", "http://localhost:8080/addemployee", bytes.NewReader(empbody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostEmployeeHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//fmt.Println(rr.Body.String())
	//expected := `{"Id":"29","Name":"Meena","Phone":"6239865088","Dept":{"Id":"121","Name":""}}`
	//if strings.TrimSpace(rr.Body.String()) != expected {
	//	t.Errorf("handler returned unexpected body: got (%v) want (%v)",
	//		rr.Body.String(), expected)
	//}
}
