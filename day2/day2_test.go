package day2

import (
	"fmt"
	"testing"

	"github.com/DomFaryna/adventOfCode2021/tools"
)

func TestPart1(t *testing.T) {
	scanner := tools.CreateScanner("input.txt")
	hor, vertical := day2_part1(scanner)
	ans := hor * vertical
	fmt.Println(ans)
}

func TestPart2(t *testing.T) {
	scanner := tools.CreateScanner("input.txt")
	hor, vertical := day2_part2(scanner)
	ans := hor * vertical
	fmt.Println(ans)
}
