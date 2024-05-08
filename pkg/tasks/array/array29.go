package array

import "fmt"

func Array29() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	max := ar[0]
	for i := 2; i < n; i += 2 {
		if max < ar[i] {
			max = ar[i]
		}
	}

	fmt.Println(max)
}
