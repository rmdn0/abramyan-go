package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	sum := 0
	for a != 0 {
		if a > 0 && a % 2 == 0 {
			sum += a
		}
		fmt.Scan(&a)
	}
	fmt.Printf("\n%d\n", sum)
}