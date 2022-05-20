package sort

import "github.com/AI1411/go-algorithm/constraints"

func Count[T constraints.Number](data []int) []int {
	var aMin, aMax = -1000, 1000
	count := make([]int, aMax-aMin+1)
	for _, x := range data {
		count[x-aMin]++
	}
	z := 0
	for i, c := range count {
		for c > 0 {
			data[z] = i + aMin
			z++
			c--
		}
	}
	return data
}
