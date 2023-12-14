package calculator

import (
	"calculator_golang/mathops"
	"fmt"
)

func Addition(a, b int) int {
	result := mathops.Add(a, b)

	return result
}

func Substarction(a, b int) int {
	result := mathops.Subtract(a, b)

	return result
}

func Multiply(a, b int) int {
	result := mathops.Multiply(a, b)

	return result
}

func PowerOf(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result = mathops.Multiply(result, base)
	}
	return result
}

func Division(numerator, denominator int) int {
	result, err := mathops.Divide(numerator, denominator)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of division:", result)
	}
	return result
}
