package internal

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
	remSpc := strings.TrimSpace(str)
	var res float64
	if res, err = strconv.ParseFloat(remSpc, 64); err != nil {
		return 0, fmt.Errorf("please, enter a numeric value")
	}
	return res, nil
}

func ReadOperation(reader *bufio.Reader) (byte, error) {
	op, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	op = strings.TrimSpace(op)

	if len(op) != 1 || !strings.Contains("+/*/", string(op[0])) {
		return 0, fmt.Errorf("please, use symbols +,-,* or /")
	}

	return op[0], nil
}
