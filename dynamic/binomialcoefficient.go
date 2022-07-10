package dynamic

import "github.com/AI1411/go-algorithm/math/min"

func Bin2(n int, k int) int {
	var i, j int
	B := make([][]int, (n + 1))
	for i := range B {
		B[i] = make([]int, k+1)
	}

	for i = 0; i <= n; i++ {
		for j = 0; j <= min.Int(i, k); j++ {
			if j == 0 || j == i {
				B[i][j] = 1
			} else {
				B[i][j] = B[i-1][j-1] + B[i-1][j]
			}
		}
	}
	return B[n][k]
}
