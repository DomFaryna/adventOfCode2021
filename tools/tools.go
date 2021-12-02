package tools

import (
	"bufio"
	"os"
)

// CreateScanner creates a new scanner
func CreateScanner(fileName string) *bufio.Scanner {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	return s
}
