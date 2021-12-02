package day1

import (
	"bufio"
	"strconv"
)

func day1(scanner *bufio.Scanner) int {
	count := 0

	scanner.Scan()
	prev, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	for scanner.Scan() {
		next, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		if next > prev {
			count++
		}
		prev = next
	}
	return count
}

func sum(i []int) int {
	count := 0
	for _, v := range i {
		count += v
	}
	return count
}

func day1_part2(scanner *bufio.Scanner) int {
	count := 0
	prev := []int{}
	for i := 0; i < 3; i++ {
		scanner.Scan()
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		prev = append(prev, val)
	}
	for scanner.Scan() {
		next := prev[1:]
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		next = append(next, val)
		if sum(next) > sum(prev) {
			count++
		}
		prev = next
	}
	return count
}
