package matrix

func Matrix3(m, n uint, floats []float64) [][]float64 {
	matrix := NewFloat64Matrix(m, n)

	for i := 0; uint(i) < m; i++ {
		for j := 0; uint(j) < n; j++ {
			matrix[i][j] = floats[i]
		}
	}

	return matrix
}
