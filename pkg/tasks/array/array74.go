package array

import "fmt"

func Array74() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	min := ar[0]
	max := ar[0]

	minID := 0
	maxID := 0

	for i := 1; i < n; i++ {
		if ar[i] < min {
			min = ar[i]
			minID = i
		}
		if ar[i] > max {
			max = ar[i]
			maxID = i
		}
	}

	start, finish := minMaxOrder(minID, maxID)

	for i := start + 1; i < finish; i++ {
		ar[i] = 0
	}

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}

func minMaxOrder(x, y int) (int, int) {
	if x > y {
		x, y = y, x
	}
	return x, y
}
