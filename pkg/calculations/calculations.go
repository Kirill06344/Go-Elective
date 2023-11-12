package calculations

import (
	log "github.com/sirupsen/logrus"
	"io"
)

func Calculate(n int64, flag bool) int64 {
	if !flag {
		log.SetOutput(io.Discard)
	}
	log.Println("Start calculations...")
	log.Printf("Calculate %d!", n)
	res := int64(1)
	for i := int64(2); i <= n; i++ {
		res *= i
	}
	log.Println("Calculations complete!")
	return res
}
