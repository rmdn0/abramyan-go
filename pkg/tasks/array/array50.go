package array

import "fmt"

func Array50() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	amount := 0
	for i := 0; i < n-1; i++ {
		if ar[i] == 1 {
			continue
		}
		for j := i + 1; j < n; j++ {
			if ar[i] > ar[j] {
				amount++
			}
		}
	}

	fmt.Println(amount)
}
