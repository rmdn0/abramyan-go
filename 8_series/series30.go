package main

import "fmt"

func main() {
	var k, n int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, sum int
	for i := 0; i < k; i++ {
		sum = 0;
		for j := 0; j < n; j++ {
			fmt.Scan(&a)
			sum += a
		}
		fmt.Printf("%d\n", sum)		
	}
}