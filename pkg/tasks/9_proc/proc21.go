package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	fmt.Printf("\n%d\n%d\n", sumRange(a, b), sumRange(b, c))
}

func sumRange(x, y int) int {
	sum := 0
	for x <= y {
		sum += x
		x++
	}
	return sum
}