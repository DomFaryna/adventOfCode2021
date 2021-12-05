package day4

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type bingoNum struct {
	verticalCount   *int
	horizontalCount *int
	called          bool
	val             int
}

type bingoCard struct {
	nums  [5][5]*bingoNum
	quick map[int]*bingoNum
}

func newBingoCard(s []string) *bingoCard {
	retVal := bingoCard{
		nums:  [5][5]*bingoNum{},
		quick: map[int]*bingoNum{},
	}
	verticals := []*int{}
	horizontals := []*int{}
	for i := 0; i < 5; i++ {
		vert := 0
		hor := 0
		verticals = append(verticals, &vert)
		horizontals = append(horizontals, &hor)
	}

	r := regexp.MustCompile("[^\\s]+")
	for horz, v := range s {
		indiv := r.FindAllString(v, -1)
		for vert, v2 := range indiv {
			num, err := strconv.Atoi(v2)
			if err != nil {
				panic(err)
			}
			bin := &bingoNum{
				called:          false,
				val:             num,
				verticalCount:   verticals[vert],
				horizontalCount: horizontals[horz],
			}
			retVal.nums[vert][horz] = bin
			retVal.quick[num] = bin
		}
	}
	return &retVal
}

func (b *bingoCard) sum() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.nums[i][j].called {
				sum += b.nums[i][j].val
			}
		}
	}
	return sum
}

func (b *bingoCard) addNum(num int) int {
	bingoNum, ok := b.quick[num]
	// number was not found on the current bingo card
	if !ok {
		return -1
	}
	bingoNum.called = true
	*bingoNum.horizontalCount = *bingoNum.horizontalCount + 1
	*bingoNum.verticalCount = *bingoNum.verticalCount + 1
	if *bingoNum.horizontalCount == 5 || *bingoNum.verticalCount == 5 {
		return b.sum() * num
	}
	return -1
}

func lineToInt(s string) []int {
	strArr := strings.Split(s, ",")
	intArr := []int{}
	for _, val := range strArr {
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		intArr = append(intArr, i)
	}
	return intArr
}

func day4_part1(s *bufio.Scanner) int {
	s.Scan()
	pickems := lineToInt(s.Text())

	bingoCards := []*bingoCard{}
	lines := []string{}
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		lines = append(lines, s.Text())
		if len(lines) == 5 {
			bingoCards = append(bingoCards, newBingoCard(lines))
			lines = []string{}
		}
	}
	for _, i := range pickems {
		for _, card := range bingoCards {
			num := card.addNum(i)
			if num != -1 {
				return num
			}
		}
	}
	fmt.Println("all done")
	return -1
}

func day4_part2(s *bufio.Scanner) int {
	s.Scan()
	pickems := lineToInt(s.Text())

	bingoCards := map[*bingoCard]struct{}{}
	lines := []string{}
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		lines = append(lines, s.Text())
		if len(lines) == 5 {
			card := newBingoCard(lines)
			bingoCards[card] = struct{}{}
			lines = []string{}
		}
	}
	for _, i := range pickems {
		for card := range bingoCards {
			num := card.addNum(i)
			if num != -1 {
				if len(bingoCards) == 1 {
					return num
				}
				delete(bingoCards, card)
			}
		}
	}
	fmt.Println("all done")
	return -1
}
