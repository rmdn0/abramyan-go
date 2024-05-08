package main

import "fmt"

func main() {
	var n, a, temp, k int
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Println()
	fmt.Scan(&a)
	temp = a
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if temp < a {
			fmt.Printf("%d ", temp)
			k++
		}
		temp = a
	}
	fmt.Printf("\nK = %d\n", k)
}