package calculator

import "fmt"

func MakeACalculation(a, b float64, op byte) (float64, error) {
	var result float64
	switch op {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		if b == 0 {
			return 0, fmt.Errorf("division by zero is impossible")
		}
		result = a / b
	default:
		return 0, fmt.Errorf("operation not yet available")
	}
	return result, nil
}
