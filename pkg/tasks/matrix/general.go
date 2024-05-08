package matrix

func NewIntMatrix(m, n uint) [][]int {
	matrix := make([][]int, m)
	for i := 0; uint(i) < m; i++ {
		for j := 0; uint(j) < n; j++ {
			matrix[i] = make([]int, n)
		}
	}

	return matrix
}
