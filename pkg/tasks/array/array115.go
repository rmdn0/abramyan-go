package array

import "fmt"

func Array115() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	indexes := make([]int, n, n)
	for i := 0; i < n; i++ {
		indexes[i] = i + 1
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if ar[indexes[j]-1] > ar[indexes[j+1]-1] {
				swapInt(&indexes[j], &indexes[j+1])
			}
		}
	}

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(indexes[i], " ")
	}
	fmt.Println()
}

func swapInt(x, y *int) {
	*x, *y = *y, *x
}
