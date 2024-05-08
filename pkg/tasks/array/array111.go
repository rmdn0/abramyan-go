package array

import (
	"fmt"
	"math"
)

func Array111() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	numberOfOdds := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
		if ar[i]%2 != 0 {
			numberOfOdds++
		}
	}

	m := n + numberOfOdds*2
	tempAr := make([]int, m, m)
	k := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= int(math.Abs(float64(ar[i]%2)))*2; j++ {
			tempAr[k] = ar[i]
			k++
		}
	}

	ar = tempAr

	fmt.Println()
	for i := 0; i < m; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}
