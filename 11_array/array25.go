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
	ratio := ar[1] / ar[0]
	for i := 2; i < n; i++ {
		if ar[i]/ar[i-1] != ratio {
			ratio = 0
			break
		}
	}

	fmt.Println(ratio)
}
