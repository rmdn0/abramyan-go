package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var k, l, m int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("L = ")
	fmt.Scan(&l)
	fmt.Print("M = ")
	fmt.Scan(&m)
	fmt.Printf("\n%.2f\n", power2(a, k))
	fmt.Printf("%.2f\n", power2(a, l))
	fmt.Printf("%.2f\n", power2(a, m))
}

func power2(a float64, n int) float64 {
	power := 1.0
	temp_n := int(math.Abs(float64(n)))
	for i := 0; i < temp_n; i++ {
		power *= a
	}
	if n < 0 {
		power = 1.0 / power
	}
	return power
}