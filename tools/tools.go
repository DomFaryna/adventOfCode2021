package tools

import (
	"bufio"
	"os"
	"strconv"
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

func StringsToInts(s []string) []int {
	retVal := []int{}
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		retVal = append(retVal, i)
	}
	return retVal
}
