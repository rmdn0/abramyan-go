package matrix

import (
	"fmt"
	"testing"
)

func TestMatrix4(t *testing.T) {
	tests := []struct {
		m               uint
		n               uint
		floats          []float64
		expectingResult [][]float64
	}{
		{
			m:      5,
			n:      6,
			floats: []float64{5.64, 5.41, 8.52, 4.88, 5.34, 3.72},
			expectingResult: [][]float64{
				{5.64, 5.41, 8.52, 4.88, 5.34, 3.72},
				{5.64, 5.41, 8.52, 4.88, 5.34, 3.72},
				{5.64, 5.41, 8.52, 4.88, 5.34, 3.72},
				{5.64, 5.41, 8.52, 4.88, 5.34, 3.72},
				{5.64, 5.41, 8.52, 4.88, 5.34, 3.72},
			},
		},
		{
			m:      3,
			n:      4,
			floats: []float64{9.36, 0.90, 8.75, 9.61},
			expectingResult: [][]float64{
				{9.36, 0.90, 8.75, 9.61},
				{9.36, 0.90, 8.75, 9.61},
				{9.36, 0.90, 8.75, 9.61},
			},
		},
		{
			m:      4,
			n:      7,
			floats: []float64{8.16, 1.06, 3.78, 7.85, 4.45, 6.32, 3.33},
			expectingResult: [][]float64{
				{8.16, 1.06, 3.78, 7.85, 4.45, 6.32, 3.33},
				{8.16, 1.06, 3.78, 7.85, 4.45, 6.32, 3.33},
				{8.16, 1.06, 3.78, 7.85, 4.45, 6.32, 3.33},
				{8.16, 1.06, 3.78, 7.85, 4.45, 6.32, 3.33},
			},
		},
	}

	for testNumber, test := range tests {
		result := Matrix4(test.m, test.n, test.floats)

		if len(result) != len(test.expectingResult) {
			t.Fatalf("test %d failed: got rows: %d; want rows: %v", testNumber+1, len(result), len(test.expectingResult))
		}

		for i := 0; uint(i) < test.m; i++ {
			if len(result[i]) != len(test.expectingResult[i]) {
				t.Fatalf("test %d failed: got columns: %d; want columns: %v", testNumber+1, len(result[i]), len(test.expectingResult[i]))
			}
			for j := 0; uint(j) < test.n; j++ {
				if result[i][j] != test.expectingResult[i][j] {
					t.Fatalf("test %d failed: got: %v; want: %v", testNumber+1, result, test.expectingResult)
				}
			}
		}

		fmt.Println("res before:", result)
		test.floats[0] = 0
		fmt.Println("res after: ", result)
	}

}
