package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	var isSetFirstEven bool
	var firstEvenIndex int
	ar := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
		if isEven(ar[i]) && !isSetFirstEven {
			firstEvenIndex = i
			isSetFirstEven = true
		}
	}

	if isSetFirstEven {
		sumKthValueToEvenElements(ar, firstEvenIndex)
	}

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}

func sumKthValueToEvenElements(ar []int, k int) {
	diff := ar[k]
	for i := 0; i < len(ar); i++ {
		if isEven(ar[i]) {
			ar[i] += diff
		}
	}
}

func isEven(x int) bool {
	return x%2 == 0
}
