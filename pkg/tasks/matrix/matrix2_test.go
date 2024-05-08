package matrix

import (
	"testing"
)

func TestMatrix2(t *testing.T) {
	tests := []struct {
		m               uint
		n               uint
		expectingResult [][]int
	}{
		{
			m: 5,
			n: 7,
			expectingResult: [][]int{
				{5, 10, 15, 20, 25, 30, 35},
				{5, 10, 15, 20, 25, 30, 35},
				{5, 10, 15, 20, 25, 30, 35},
				{5, 10, 15, 20, 25, 30, 35},
				{5, 10, 15, 20, 25, 30, 35},
			},
		},
		{
			m: 4,
			n: 6,
			expectingResult: [][]int{
				{5, 10, 15, 20, 25, 30},
				{5, 10, 15, 20, 25, 30},
				{5, 10, 15, 20, 25, 30},
				{5, 10, 15, 20, 25, 30},
			},
		},
		{
			m: 3,
			n: 9,
			expectingResult: [][]int{
				{5, 10, 15, 20, 25, 30, 35, 40, 45},
				{5, 10, 15, 20, 25, 30, 35, 40, 45},
				{5, 10, 15, 20, 25, 30, 35, 40, 45},
			},
		},
	}

	for testNumber, test := range tests {
		result := Matrix2(test.m, test.n)

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
	}

}
