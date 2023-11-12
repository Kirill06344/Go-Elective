package main

import (
	"GoBasics/pkg/calculations"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	usingLogger := flag.Bool("log", false, "using logger")
	flag.Parse()
	if len(flag.Args()) != 1 {
		return
	}

	n, err := strconv.ParseInt(flag.Args()[0], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	answer := calculations.Calculate(n, *usingLogger)
	fmt.Printf("Factorial of %d: %d \n", n, answer)

}
