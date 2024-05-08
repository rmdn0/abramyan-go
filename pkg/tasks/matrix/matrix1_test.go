package matrix

import (
	"testing"
)

func TestMatrix1(t *testing.T) {
	tests := []struct {
		m               uint
		n               uint
		expectingResult [][]int
	}{
		{
			m: 5,
			n: 3,
			expectingResult: [][]int{
				{10, 10, 10},
				{20, 20, 20},
				{30, 30, 30},
				{40, 40, 40},
				{50, 50, 50},
			},
		},
		{
			m: 4,
			n: 8,
			expectingResult: [][]int{
				{10, 10, 10, 10, 10, 10, 10, 10},
				{20, 20, 20, 20, 20, 20, 20, 20},
				{30, 30, 30, 30, 30, 30, 30, 30},
				{40, 40, 40, 40, 40, 40, 40, 40},
			},
		},
		{
			m: 5,
			n: 6,
			expectingResult: [][]int{
				{10, 10, 10, 10, 10, 10},
				{20, 20, 20, 20, 20, 20},
				{30, 30, 30, 30, 30, 30},
				{40, 40, 40, 40, 40, 40},
				{50, 50, 50, 50, 50, 50},
			},
		},
	}

	for testNumber, test := range tests {
		result := Matrix1(test.m, test.n)

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
