package day3

import (
	"fmt"
	"testing"

	"github.com/DomFaryna/adventOfCode2021/tools"
)

func TestPart1(t *testing.T) {
	scanner := tools.CreateScanner("input.txt")
	ans := day3_part1(scanner, 12)
	fmt.Printf("The Answer is %d \n", ans)
}

func TestPart2(t *testing.T) {
	scanner := tools.CreateScanner("input.txt")
	ans := day3_part2(scanner)
	fmt.Printf("The Answer is %d \n", ans)
}
