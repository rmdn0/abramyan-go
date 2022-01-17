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
	fmt.Printf("\nGCD3 (A, B, C) = %d\n", gcd3(a, b, c))
	fmt.Printf("GCD3 (A, C, D) = %d\n", gcd3(a, c, d))
	fmt.Printf("GCD3 (B, C, D) = %d\n", gcd3(b, c, d))
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

func gcd3(a, b, c int) int {
	return gcd2(gcd2(a, b), c)
}