package matrix

func Matrix4(m, n uint, floats []float64) [][]float64 {
	matrix := NewFloat64Matrix(m, n)

	for i := 0; uint(i) < m; i++ {
		//matrix[i] = floats // will copy the pointer
		for j := 0; uint(j) < n; j++ {
			matrix[i][j] = floats[j]
		}
	}

	return matrix
}
