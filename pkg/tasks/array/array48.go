package array

import "fmt"

func Array48() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	var amount, maxAmount int
	for i := 0; i < n; i++ {
		amount = 0
		for j := 0; j < n; j++ {
			if ar[i] == ar[j] {
				amount++
			}
		}
		if maxAmount < amount {
			maxAmount = amount
		}
	}

	fmt.Println(maxAmount)
}
