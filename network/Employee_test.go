package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type test struct {
	method string
	url    string
	output interface{}
	er     error
}

var tests = []test{
	{"GET", "/emp?id=11", true, nil},
	{"GET", "/emp?id=13", false, errors.New("invalid Id")},
	{"POST", "/emp", true, nil},
}

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/emp?Id=12", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Employee)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":"12","name":"Saikiran","age":"22","address":"Hyderabad"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got (%v) want (%v)",
			rr.Body.String(), expected)
	}
}
