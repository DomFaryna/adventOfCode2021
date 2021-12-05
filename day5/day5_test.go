package day5

import (
	"fmt"
	"testing"

	"github.com/DomFaryna/adventOfCode2021/tools"
)

func TestPart1(t *testing.T) {
	scanner := tools.CreateScanner("input.txt")
	num := day5_part1(scanner)
	fmt.Println(num)
}

func TestPart2(t *testing.T) {
	scanner := tools.CreateScanner("input.txt")
	num := day5_part2(scanner)
	fmt.Println(num)
}
