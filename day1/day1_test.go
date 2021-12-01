package day1

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)
	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	total := day1(s)
	fmt.Println(total)
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)
	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	total := day1_part2(s)
	fmt.Println(total)
}
