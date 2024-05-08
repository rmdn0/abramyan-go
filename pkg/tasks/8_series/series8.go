package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, k int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a % 2 == 0 {
			fmt.Println(a)
			k++
		}
	}
	fmt.Printf("\nK = %d\n", k)
}