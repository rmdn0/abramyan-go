package array

import "fmt"

func Array54() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	evenCounter := 0
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if isEven(a[i]) {
			evenCounter++
		}
	}

	b := make([]int, evenCounter, evenCounter)
	var j int
	for i := 0; i < n; i++ {
		if isEven(a[i]) {
			b[j] = a[i]
			j++
		}
	}

	fmt.Println("\nSize of array B:", evenCounter)
	for i := 0; i < evenCounter; i++ {
		fmt.Print(b[i], " ")
	}
	fmt.Println()
}
