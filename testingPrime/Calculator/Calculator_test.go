package Calculator

import (
	"errors"
	"testing"
)

func TestAddition(t *testing.T) {
	type additiontest struct {
		val1 int
		val2 int
		res  int
	}
	additiontests := []additiontest{
		{2, 3, 5},
		{-5, 5, 0},
		{5, -10, -5},
		{-10, -4, -14},
	}
	for _, num := range additiontests {
		got := Add(num.val1, num.val2)
		expect := num.res
		if got != expect {
			t.Errorf("got : %v Expected : %v", got, expect)
		}
	}
}

func TestSub(t *testing.T) {
	type subtracttest struct {
		val1 int
		val2 int
		res  int
	}
	subtracttests := []subtracttest{
		{14, 3, 11},
		{15, 17, -2},
		{-5, -10, 5},
		{-9, -4, -5},
	}
	for _, num := range subtracttests {
		got := Sub(num.val1, num.val2)
		expect := num.res
		if got != expect {
			t.Errorf("got : %v Expected : %v", got, expect)
		}
	}
}

func TestMultiply(t *testing.T) {
	type multiplytest struct {
		val1 int
		val2 int
		res  int
	}
	multiplytests := []multiplytest{
		{2, 3, 6},
		{-5, 5, -25},
		{5, -10, -50},
		{-10, -4, 40},
	}
	for _, num := range multiplytests {
		got := Multiply(num.val1, num.val2)
		expect := num.res
		if got != expect {
			t.Errorf("got : %v Expected : %v", got, expect)
		}
	}
}

func TestDivide(t *testing.T) {
	type dividetest struct {
		val1 int
		val2 int
		res  int
		er   error
	}
	dividetests := []dividetest{
		{4, 2, 2, nil},
		{5, 2, 2, nil},
		{-4, -2, 2, nil},
		{-4, 2, -2, nil},
		{5, 0, 0, errors.New("cannot divide by zero")},
		{0, 23, 0, nil},
	}
	for _, val := range dividetests {
		got, err := Divide(val.val1, val.val2)
		expect := val.res
		if err != nil && err.Error() != val.er.Error() {
			t.Errorf("error : %v", err)
			return
		}
		if got != expect {
			t.Errorf("got : %v Expected : %v", got, expect)
		}
	}
}

func TestModulo(t *testing.T) {
	type modulotest struct {
		val1 int
		val2 int
		res  int
		er   error
	}
	modulotests := []modulotest{
		{4, 2, 0, nil},
		{5, 2, 1, nil},
		{-4, -2, 0, nil},
		{-3, 2, -1, nil},
		{5, 0, 0, errors.New("cannot modulo by zero")}, // Error is evaluated at line 115
		{0, 23, 0, nil},
	}
	for _, val := range modulotests {
		got, err := Modulo(val.val1, val.val2)
		expect := val.res
		if err != nil && err.Error() != val.er.Error() {
			t.Errorf("error : %v", err)
			return
		}
		if got != expect {
			t.Errorf("got : %v Expected : %v", got, expect)
		}
	}
}

func TestOperation(t *testing.T) {
	type Optest struct {
		val1 int
		val2 int
		op   string
		res  int
		er   error
	}
	tests := []Optest{
		{2, 3, "+", 5, nil},
		{4, 6, "-", -2, nil},
		{5, 3, "*", 15, nil},
		{4, 2, "/", 2, nil},
		{7, 2, "/", 3, nil},
		{13, 0, "/", 0, errors.New("cannot divide by zero")},
		{12, 10, "%", 2, nil},
		{20, 10, "%", 0, nil},
		{7, 0, "%", 0, errors.New("cannot modulo by zero")},
		{12, 6, ")", -1, errors.New("invalid operator")},
	}
	for _, val := range tests {
		got, err := Operation(val.val1, val.val2, val.op)
		expect := val.res
		if err != nil && err.Error() != val.er.Error() {
			t.Errorf("Error : %v", err)
		}
		if got != expect {
			t.Errorf("got : %v Expected : %v", got, expect)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(300, 1276)
	}
}

func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sub(3086, 12)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(300, 1276)
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(300, 14)
	}
}

func BenchmarkModulo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Modulo(234, 75)
	}
}

func BenchmarkOperation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Operation(23, 21, "*")
	}
}
