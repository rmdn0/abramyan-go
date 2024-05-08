package array

import "fmt"

func Array110() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	numberOfEvens := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
		if ar[i]%2 == 0 {
			numberOfEvens++
		}
	}

	m := n + numberOfEvens
	tempAr := make([]int, m, m)
	k := 0
	for i := 0; i < n; i++ {
		tempAr[k] = ar[i]
		k++
		if ar[i]%2 == 0 {
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
