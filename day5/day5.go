package day5

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/DomFaryna/adventOfCode2021/tools"
)

type point struct {
	x int
	y int
}

type line struct {
	start *point
	end   *point
}

func (l *line) getPoints(orthOnly bool) []*point {
	points := []*point{}
	if l.start.x == l.end.x {
		if l.start.y > l.end.y {
			l.start.y, l.end.y = l.end.y, l.start.y
		}
		for i := l.start.y; i <= l.end.y; i++ {
			p := point{
				x: l.start.x,
				y: i,
			}
			points = append(points, &p)
		}
		return points
	}
	if l.start.y == l.end.y {
		if l.start.x > l.end.x {
			l.start.x, l.end.x = l.end.x, l.start.x
		}
		for i := l.start.x; i <= l.end.x; i++ {
			p := point{
				x: i,
				y: l.start.y,
			}
			points = append(points, &p)
		}
		return points
	}

	if !orthOnly {
		xIncrement := 1
		if l.start.x > l.end.x {
			xIncrement = -1
		}
		yIncrement := 1
		if l.start.y > l.end.y {
			yIncrement = -1
		}
		x := l.start.x
		y := l.start.y
		for x != l.end.x {
			p := point{
				x: x,
				y: y,
			}
			points = append(points, &p)
			x += xIncrement
			y += yIncrement
		}
		p := point{
			x: l.end.x,
			y: l.end.y,
		}
		points = append(points, &p)
	}

	return points
}

func newLine(x1, y1, x2, y2 int) *line {
	startPoint := point{x: x1, y: y1}
	endPoint := point{x: x2, y: y2}
	l := line{
		start: &startPoint,
		end:   &endPoint,
	}
	return &l
}

type grid [1000][1000]int

func newGrid() grid {
	g := grid{}
	return g
}

func (g *grid) addLine(l *line, orthOnly bool) {
	for _, p := range l.getPoints(orthOnly) {
		g[p.y][p.x]++
	}
}

func (g *grid) getOverlap() int {
	sum := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if g[i][j] >= 2 {
				sum++
			}
		}
	}
	return sum
}

func (g *grid) debugPrint() {
	for i := 0; i < 10; i++ {
		sb := strings.Builder{}
		for j := 0; j < 10; j++ {
			if g[i][j] != 0 {
				fmt.Fprintf(&sb, "%d", g[i][j])
				continue
			}
			fmt.Fprintf(&sb, ".")
		}
		fmt.Println(sb.String())
	}
}

func day5_part1(s *bufio.Scanner) int {
	g := newGrid()
	r := regexp.MustCompile("[0-9]+")
	for s.Scan() {
		nums := tools.StringsToInts(r.FindAllString(s.Text(), -1))
		l := newLine(nums[0], nums[1], nums[2], nums[3])

		g.addLine(l, true)
	}
	count := g.getOverlap()
	g.debugPrint()
	return count
}

func day5_part2(s *bufio.Scanner) int {
	g := newGrid()
	r := regexp.MustCompile("[0-9]+")
	for s.Scan() {
		nums := tools.StringsToInts(r.FindAllString(s.Text(), -1))
		l := newLine(nums[0], nums[1], nums[2], nums[3])
		g.addLine(l, false)
	}
	count := g.getOverlap()
	g.debugPrint()
	return count
}
