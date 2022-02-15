package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	ar := make([]int, n, n)
	powerOfTwo := 1
	for i := 0; i < n; i++ {
		powerOfTwo *= 2
		ar[i] = powerOfTwo
	}
	printArray(ar)
}

func printArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}
