package day4

import (
	"fmt"
	"testing"

	"github.com/DomFaryna/adventOfCode2021/tools"
)

func TestBingo(t *testing.T) {
	sample := []string{}
	sample = append(sample, "14 21 17 24  4")
	sample = append(sample, "10 16 15  9 19")
	sample = append(sample, "18  8 23 26 20")
	sample = append(sample, "22 11 13  6  5")
	sample = append(sample, " 2  0 12  3  7")

	card := newBingoCard(sample)
	card.addNum(14)
	card.addNum(21)
	card.addNum(17)
	card.addNum(4)
	card.addNum(9)
	card.addNum(23)
	card.addNum(11)
	card.addNum(5)
	card.addNum(2)
	card.addNum(0)
	card.addNum(7)
	num := card.addNum(24)

	fmt.Println(num)
}

func TestPart1(t *testing.T) {
	scanner := tools.CreateScanner("input.txt")
	num := day4_part1(scanner)
	fmt.Println(num)
}

func TestPart2(t *testing.T) {
	scanner := tools.CreateScanner("input.txt")
	num := day4_part2(scanner)
	fmt.Println(num)
}
