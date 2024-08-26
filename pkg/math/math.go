package math

import (
	"errors"
)

// Add Functioner
func AddValues(x, y int) int {
	return x + y
}

// Func for division
func DivideValues(x, y float64) (float64, error) {
	if y == 0 {
		err := errors.New("Cannot divide by zero!")
		return 0, err
	}
	return x / y, nil
}
