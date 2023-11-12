package main

import (
	"GoBasics/pkg/calculations"

	log "github.com/sirupsen/logrus"

	"fmt"
	"os"
	"strconv"
)

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("Start calculations...")
	log.Printf("Calculate %d!", n)
	answer := calculations.Calculate(n, false)
	fmt.Println(answer)
}
