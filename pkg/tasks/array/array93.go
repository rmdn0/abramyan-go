package array

import "fmt"

func Array93() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	m := n/2 + n%2
	tempAr := make([]int, m, m)
	j := 0
	for i := 0; i < n; i += 2 {
		tempAr[j] = ar[i]
		j++
	}

	ar = tempAr
	fmt.Println()
	for i := 0; i < m; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}
