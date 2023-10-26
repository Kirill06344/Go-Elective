package main

import (
	"GoBasics/pkg/calculator"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first operand: ")
	a, err := ReadFloat64(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Select operation(+,-,*,/): ")
	op, err := ReadOperation(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter second operand: ")
	b, err := ReadFloat64(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := calculator.MakeACalculation(a, b, op)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%g %s %g = %g", a, string(op), b, result)
}
