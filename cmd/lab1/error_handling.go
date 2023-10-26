package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ReadFloat64(reader *bufio.Reader) (float64, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	remSpc := str[:len(str)-1]
	var res float64
	if res, err = strconv.ParseFloat(remSpc, 64); err != nil {
		return 0, fmt.Errorf("please, enter a numeric value")
	}
	return res, nil
}

func ReadOperation(reader *bufio.Reader) (byte, error) {
	op, err := reader.ReadBytes('\n')
	if err != nil {
		return 0, err
	}

	if len(op) != 2 || !strings.Contains("+/*/", string(op[0])) {
		return 0, fmt.Errorf("please, use symbols +,-,* or /")
	}

	return op[0], nil
}
