package day6

func cycle(ar []int, days int) int {
	a := []int{}

	for i := 0; i < 9; i++ {
		a = append(a, 0)
	}
	for _, v := range ar {
		a[v]++
	}
	for i := 0; i < days; i++ {
		zeros := a[0]
		for j := 0; j < 8; j++ {
			a[j] = a[j+1]
		}
		a[6] += zeros
		a[8] = zeros
	}
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
