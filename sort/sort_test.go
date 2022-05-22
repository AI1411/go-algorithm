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

func TestCombSort(t *testing.T) {
	testFramework(t, sort.Comb[int])
}

func TestCount(t *testing.T) {
	testFramework(t, sort.Count[int])
}

func TestExchange(t *testing.T) {
	testFramework(t, sort.Exchange[int])
}

func TestHeap(t *testing.T) {
	testFramework(t, sort.HeapSort)
}

func TestInsertionSort(t *testing.T) {
	testFramework(t, sort.InsertionSort[int])
}

func TestMerge(t *testing.T) {
	testFramework(t, sort.Merge[int])
}

func benchmarkFramework(b *testing.B, f func(arr []int) []int) {
	var sortTests = []struct {
		input    []int
		expected []int
		name     string
	}{
		//Sorted slice
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Sorted Unsigned"},
		//Reversed slice
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed Unsigned"},
		//Sorted slice
		{[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Sorted Signed"},
		//Reversed slice
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed Signed"},
		//Reversed slice, even length
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed Signed #2"},
		//Random order with repetitions
		{[]int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10}, "Random order Signed"},
		//Empty slice
		{[]int{}, []int{}, "Empty"},
		//Single-entry slice
		{[]int{1}, []int{1}, "Singleton"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			f(test.input)
		}
	}
}

func BenchmarkBubble(b *testing.B) {
	benchmarkFramework(b, sort.Bubble[int])
}

func BenchmarkComb(b *testing.B) {
	benchmarkFramework(b, sort.Comb[int])
}

func BenchmarkCount(b *testing.B) {
	benchmarkFramework(b, sort.Count[int])
}

func BenchmarkExchange(b *testing.B) {
	benchmarkFramework(b, sort.Exchange[int])
}

func BenchmarkHeap(b *testing.B) {
	benchmarkFramework(b, sort.HeapSort)
}

func BenchmarkInsertion(b *testing.B) {
	benchmarkFramework(b, sort.InsertionSort[int])
}

func BenchmarkMerge(b *testing.B) {
	benchmarkFramework(b, sort.Merge[int])
}
