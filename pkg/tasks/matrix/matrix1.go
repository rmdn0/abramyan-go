package matrix

func Matrix1(m, n uint) [][]int {
	matrix := make([][]int, m, m)
	for i := 0; uint(i) < m; i++ {
		for j := 0; uint(j) < n; j++ {
			matrix[i] = make([]int, n, n)
		}
	}

	for i := 0; uint(i) < m; i++ {
		for j := 0; uint(j) < n; j++ {
			matrix[i][j] = (i + 1) * 10
		}
	}

	return matrix
}
