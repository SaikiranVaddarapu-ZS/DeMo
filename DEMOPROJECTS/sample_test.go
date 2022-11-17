package main

import "testing"

type addTest struct {
	val      int
	expected bool
}

var addTests = []addTest{
	addTest{2, true},
	addTest{6, false},
	addTest{25, false},
	addTest{31, true},
}

func TestIsPrime(t *testing.T) {
	for index := range addTests {
		got := IsPrime(addTests[index].val)
		expected := addTests[index].expected
		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	}
}
