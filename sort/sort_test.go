package sort_test

import (
	"reflect"
	"testing"

	"github.com/AI1411/go-algorithm/sort"
)

func testFramework(t *testing.T, sortingFunction func([]int) []int) {
	sortTests := []struct {
		input []int
		want  []int
		name  string
	}{
		{
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:  "Sorted Unsigned Ints",
		},
		{
			input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:  "Reverse Sorted Unsigned Ints",
		},
		{
			input: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:  "Sorted Signed Ints",
		},
		{
			input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			want:  []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:  "Reverse Sorted Signed Ints",
		},
		{
			input: []int{1},
			want:  []int{1},
			name:  "Single Element",
		},
		{
			input: []int{},
			want:  []int{},
			name:  "Empty Slice",
		},
	}
	for _, test := range sortTests {
		t.Run(test.name, func(t *testing.T) {
			got := sortingFunction(test.input)
			sorted := reflect.DeepEqual(got, test.want)
			if !sorted {
				t.Errorf("test %s failed", test.name)
				t.Errorf("actual %v want %v", got, test.want)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	testFramework(t, sort.Bubble[int])
}
