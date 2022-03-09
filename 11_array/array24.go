package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	var diff int

	if n > 1 {
		diff = ar[1] - ar[0]
	} else {
		diff = ar[9]
	}
	for i := 2; i < n; i++ {
		if ar[i]-ar[i-1] != diff {
			diff = 0
			break
		}
	}
	fmt.Println(diff)
}
