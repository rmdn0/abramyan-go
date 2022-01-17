package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	fmt.Print("D = ")
	fmt.Scan(&d)
	fmt.Printf("\nGCD2 (A, B) = %d\n", gcd2(a, b))
	fmt.Printf("GCD2 (A, C) = %d\n", gcd2(a, c))
	fmt.Printf("GCD2 (A, D) = %d\n", gcd2(a, d))
}

func gcd2(a, b int) int {
	var temp_a int
	for b != 0 {
		temp_a = a
		a = b
		b = temp_a % b
	}
	return a
}