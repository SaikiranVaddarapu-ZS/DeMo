package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Emp struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Age     string `json:"age"`
	Address string `json:"address"`
}

func (data Emp) String() string {
	return fmt.Sprintf("Id : %v, Name : %v, Age : %v, Address : %v", data.Id, data.Name, data.Age, data.Address)
}

var data []Emp

func initialData() {
	data = append(data, Emp{"12", "Saikiran", "22", "Hyderabad"})
}

func addEmployee(e Emp, data []Emp) []Emp {
	data = append(data, e)
	return data
}

func getEmployee(id string, data []Emp) (Emp, error) {
	for _, val := range data {
		if val.Id == id {
			return val, nil
		}
	}
	return Emp{}, errors.New("invalid ID")
}

func Employee(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		empID := mux.Vars(r)["Id"]
		reqData, err := getEmployee(empID, data)
		if err != nil {
			w.WriteHeader(404)
			return
		}
		json.NewEncoder(w).Encode(reqData)
		w.WriteHeader(http.StatusOK)
	case "POST":
		var emp Emp
		w.Header().Set("Content-Type", "application/json")
		req, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "enter data")
		}
		json.Unmarshal(req, &emp)
		data = addEmployee(emp, data)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	initialData()
	http.HandleFunc("/emp", Employee)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
