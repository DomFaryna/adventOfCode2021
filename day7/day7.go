package day7

import (
	"math"
	"sort"

	"gonum.org/v1/gonum/stat"
)

func day7_part1(crabs []int) int {
	sort.Ints(crabs)
	crabsFloat := []float64{}
	for _, a := range crabs {
		crabsFloat = append(crabsFloat, float64(a))
	}
	median := stat.Quantile(0.5, stat.Empirical, crabsFloat, nil)
	sum := 0
	for _, v := range crabsFloat {
		sum += int(math.Abs(v - median))
	}

	return sum
}

func day7_part2(crabs []int) float64 {
	sort.Ints(crabs)
	crabsFloat := []float64{}
	for _, a := range crabs {
		crabsFloat = append(crabsFloat, float64(a))
	}
	median := stat.Mean(crabsFloat, nil)
	avg := int(median)
	sum := float64(0)
	for _, v := range crabs {
		n := math.Abs(float64(v - avg))
		distance := (n * ((n + 1) / 2))
		sum += distance
	}

	return sum
}
