package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	min := ar[1]
	for i := 3; i < n; i += 2 {
		if min > ar[i] {
			min = ar[i]
		}
	}

	fmt.Println(min)
}
