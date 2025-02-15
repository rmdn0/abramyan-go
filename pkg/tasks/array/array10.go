package array

import "fmt"

func Array10() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	for i := 0; i < n; i++ {
		if ar[i]%2 == 0 {
			fmt.Printf("%d ", ar[i])
		}
	}

	for i := n - 1; i >= 0; i-- {
		if ar[i]%2 != 0 {
			fmt.Printf("%d ", ar[i])
		}
	}

	fmt.Println()
}
