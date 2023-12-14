package mathops

import (
	"errors"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("division by zero error")
	} else {
		return a / b, nil
	}
}
