package Calculator

import (
	"errors"
)

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func Modulo(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot modulo by zero")
	}
	return a % b, nil
}

func Operation(a int, b int, op string) (int, error) {
	switch op {
	case "+":
		return Add(a, b), nil
	case "-":
		return Sub(a, b), nil
	case "*":
		return Multiply(a, b), nil
	case "/":
		return Divide(a, b)
	case "%":
		return Modulo(a, b)
	default:
		return -1, errors.New("invalid operator")
	}
}
