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

func NewFloat64Matrix(m, n uint) [][]float64 {
	matrix := make([][]float64, m)
	for i := 0; uint(i) < m; i++ {
		for j := 0; uint(j) < n; j++ {
			matrix[i] = make([]float64, n)
		}
	}

	return matrix
}
