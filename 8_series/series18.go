package main

import "fmt"

func main() {
	var n, a, temp int
	fmt.Print("N = ")
	fmt.Scan(&n)
	if n > 0 {
		fmt.Scan(&a)
		temp = a
		fmt.Printf("\n%d ", a)
	}
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if temp == a {
			continue
		}
		fmt.Printf("%d ", a)
		temp = a
	}
	fmt.Println()
}