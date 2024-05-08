package array

import (
	"fmt"
)

func Array96() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	tempAr := make([]int, 0, n)
	uniqueElements := make(map[int]struct{})
	for i := 0; i < n; i++ {
		if _, ok := uniqueElements[ar[i]]; !ok {
			uniqueElements[ar[i]] = struct{}{}
			tempAr = append(tempAr, ar[i])
		}
	}

	ar = tempAr

	fmt.Println()
	for i := 0; i < len(ar); i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}
