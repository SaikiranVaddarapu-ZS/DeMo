package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
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
	{"GET", "/", true, nil},
}

func TestHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "localhost:8080/emp?Id=12", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Employee)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `{"id":"12","name":"Saikiran","age":"22","address":"Hyderabad"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
