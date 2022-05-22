package sort

import (
	"github.com/AI1411/go-algorithm/constraints"
	"github.com/AI1411/go-algorithm/math/min"
)

func merge[T constraints.Ordered](a []T, b []T) []T {
	r := make([]T, len(a)+len(b))
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}
	return r
}

func Merge[T constraints.Ordered](item []T) []T {
	if len(item) < 2 {
		return item
	}
	mid := len(item) / 2
	a := Merge(item[:mid])
	b := Merge(item[mid:])
	return merge(a, b)
}

func MergeIter[T constraints.Ordered](items []T) []T {
	for step := 1; step < len(items); step += step {
		for i := 0; i < len(items); i += 2 * step {
			tmp := merge(items[i:i+step], items[i+step:min.Int(i+2*step, len(items))])
			copy(items[i:], tmp)
		}
	}
	return items
}
