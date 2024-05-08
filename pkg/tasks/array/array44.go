package array

import "fmt"

func Array44() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if ar[i] == ar[j] {
				fmt.Println(i+1, j+1)
				break
			}
		}
	}
}
