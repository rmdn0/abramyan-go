package array

import "fmt"

func Array26() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	ar := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	breakingElement := 0
	for i := 1; i < n; i++ {
		if isEven(ar[i]) == isEven(ar[i-1]) {
			breakingElement = i + 1
			break
		}
	}

	fmt.Println(breakingElement)
}

func isEven(x int) bool {
	return x%2 == 0
}
