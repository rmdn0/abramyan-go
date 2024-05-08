package main

import "fmt"

func main() {
	var k, n int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, sum int
	const desired_number = 2
	var has_2 bool
	for i := 0; i < k; i++ {
		has_2 = false
		sum = 0
		for j := 0; j < n; j++ {
			fmt.Scan(&a)
			if a == desired_number && !has_2 {
				has_2 = true
			}
			sum += a
		}
		if has_2 {
			fmt.Printf("%d\n", sum)
		} else {
			fmt.Printf("%d\n", 0)
		}
	}
}