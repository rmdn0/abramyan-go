package matrix

func Matrix2(m, n uint) [][]int {
	matrix := NewIntMatrix(m, n)

	for i := 0; uint(i) < m; i++ {
		for j := 0; uint(j) < n; j++ {
			matrix[i][j] = (j + 1) * 5
		}
	}

	return matrix
}
