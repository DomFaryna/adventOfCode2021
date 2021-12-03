package day3

import (
	"bufio"
	"fmt"
)

func day3_part1(scanner *bufio.Scanner, size int) int {
	bits := []int{}
	for i := 0; i < size; i++ {
		bits = append(bits, 0)
	}
	len := 0
	for scanner.Scan() {
		len++
		line := scanner.Text()
		for i, v := range line {
			if v == '1' {
				bits[i]++
			}
		}
	}
	gamma := 0
	epilson := 0
	for _, v := range bits {
		if v >= (len / 2) {
			gamma += 1
		} else {
			epilson += 1
		}
		fmt.Printf("SubAnswer is %d\n", gamma*epilson)

		gamma = gamma << 1
		epilson = epilson << 1
	}
	gamma = gamma >> 1
	epilson = epilson >> 1
	return gamma * epilson
}

func day3_part2(scanner *bufio.Scanner) int {
	oxygenLines := []string{}
	carbonLines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		oxygenLines = append(oxygenLines, line)
		carbonLines = append(carbonLines, line)
	}
	// first find the oxygen rating
	digit := 0
	for len(oxygenLines) != 1 {
		count := 0
		for _, line := range oxygenLines {
			if line[digit] == '1' {
				count++
			}
		}
		newOxygenLines := []string{}
		comp := byte('0')
		if count*2 >= len(oxygenLines) {
			comp = byte('1')
		}
		for _, line := range oxygenLines {
			if line[digit] == comp {
				newOxygenLines = append(newOxygenLines, line)
			}
		}
		oxygenLines = newOxygenLines
		digit++
	}

	// Now do carbon
	digit = 0
	for len(carbonLines) != 1 {
		count := 0
		for _, line := range carbonLines {
			if line[digit] == '1' {
				count++
			}
		}
		newCarbonLines := []string{}
		comp := byte('0')
		if count*2 < len(carbonLines) {
			comp = byte('1')
		}
		for _, line := range carbonLines {
			if line[digit] == comp {
				newCarbonLines = append(newCarbonLines, line)
			}
		}
		carbonLines = newCarbonLines
		digit++
	}

	oxygen := stringToInt(oxygenLines[0])
	carbon := stringToInt(carbonLines[0])

	return oxygen * carbon
}

func stringToInt(s string) int {
	ret := 0
	for _, v := range s {
		if v == '1' {
			ret += 1
		}
		ret = ret << 1
	}
	ret = ret >> 1
	return ret
}
