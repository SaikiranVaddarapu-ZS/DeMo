package main

import "testing"

type sides struct {
	side1  int
	side2  int
	side3  int
	result string
}

var sidesTests = []sides{
	sides{0, 0, 0, "Sides cannot be zero"},
	sides{1, 0, 0, "Sides cannot be zero"},
	sides{10, 23, 0, "Sides cannot be zero"},
	sides{5, 7, 13, "Cannot form a Triangle"},
	sides{12, 9, 1, "Cannot form a Triangle"},
	sides{5, 5, 5, "Equilateral"},
	sides{5, 5, 6, "Isosceles"},
	sides{5, 64, 64, "Isosceles"},
	sides{5, 6, 7, "Scalene"},
}

func TestTriangle(t *testing.T) {
	for index := range sidesTests {
		got := Triangle(sidesTests[index].side1, sidesTests[index].side2, sidesTests[index].side3)
		expected := sidesTests[index].result
		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	}
}
