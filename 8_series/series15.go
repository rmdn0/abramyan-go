package main

import "fmt"

func main() {
	var k, a int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Scan(&a)
	i := 1
	a_greater_than_k_index := 0
	for a != 0 {
		if a > k && a_greater_than_k_index == 0 {
			a_greater_than_k_index = i
		}
		fmt.Scan(&a)
		i++
	}
	fmt.Printf("\n%d\n", a_greater_than_k_index)
}