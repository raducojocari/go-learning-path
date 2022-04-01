package simplemath

import (
	"errors"
)

//variadic function
func Sum(values ...float64) float64 {
	total := 0.0

	for _, value := range values {
		total += value
	}
	return total
}

func Divide(a, b float64) (answer float64, err error) {
	if b == 0.0 {
		err = errors.New("arg cannot be zero")
	}

	answer = a / b
	return
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Subtract(a, b float64) float64 {
	return a - b
}
