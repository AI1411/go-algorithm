package sort

import (
	"github.com/AI1411/go-algorithm/math/max"
	"github.com/AI1411/go-algorithm/math/min"
)

func PipeOnHole(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	max := max.Int(arr...)
	min := min.Int(arr...)

	size := max - min + 1

	holes := make([]int, size)

	for _, element := range arr {
		holes[element-min]++
	}

	i := 0

	for j := 0; j < size; j++ {
		for holes[j] > 0 {
			holes[j]--
			arr[i] = j + min
			i++
		}
	}
	return arr
}
