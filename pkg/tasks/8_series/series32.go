package main

import "fmt"

func main() {
	var k, n int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, desired_index int
	const desired_number = 2
	for i := 0; i < k; i++ {
		desired_index = 0
		for j := 0; j < n; j++ {
			fmt.Scan(&a)
			if a == desired_number && desired_index == 0 {
				desired_index = j + 1
			}
		}
		fmt.Printf("%d\n", desired_index)
	}
}