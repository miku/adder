package adder

import "fmt"

func Add(a, b int) int {
	return a + b
}

func AddErr(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a + b, nil
}
