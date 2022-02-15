package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)

	ar := make([]int, n, n)
	ar[0] = a
	ar[1] = b
	sum := a + b

	for i := 2; i < n; i++ {
		ar[i] = sum
		sum *= 2
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", ar[i])
	}
	fmt.Println()
}
