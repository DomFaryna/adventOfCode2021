package day2

import (
	"bufio"
	"strconv"
	"strings"
)

func day2_part1(scanner *bufio.Scanner) (int, int) {
	horizontal := 0
	vertical := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "forward") {
			line = strings.ReplaceAll(line, "forward ", "")
			amount, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			horizontal += amount
			continue
		}
		if strings.Contains(line, "down") {
			line = strings.ReplaceAll(line, "down ", "")
			amount, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			vertical += amount
			continue
		}
		if strings.Contains(line, "up") {
			line = strings.ReplaceAll(line, "up ", "")
			amount, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			vertical -= amount
			continue
		}
	}
	return horizontal, vertical
}

func day2_part2(scanner *bufio.Scanner) (int, int) {
	horizontal := 0
	vertical := 0
	aim := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "forward") {
			line = strings.ReplaceAll(line, "forward ", "")
			amount, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			horizontal += amount
			vertical += aim * amount
			continue
		}
		if strings.Contains(line, "down") {
			line = strings.ReplaceAll(line, "down ", "")
			amount, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			aim += amount
			continue
		}
		if strings.Contains(line, "up") {
			line = strings.ReplaceAll(line, "up ", "")
			amount, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			aim -= amount
			continue
		}
	}
	return horizontal, vertical
}
