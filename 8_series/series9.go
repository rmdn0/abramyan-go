package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, k int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a % 2 == 1 || a % 2 == -1 {
			fmt.Println(i + 1)
			k++
		}
	}
	fmt.Printf("\nK = %d\n", k)
}