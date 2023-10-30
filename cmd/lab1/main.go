package main

import (
	"bufio"
	"fmt"
	"os"

	"GoBasics/internal"
	"GoBasics/pkg/calculator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first operand: ")
	a, err := internal.ReadFloat64(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Select operation(+,-,*,/): ")
	op, err := internal.ReadOperation(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter second operand: ")
	b, err := internal.ReadFloat64(reader)
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
