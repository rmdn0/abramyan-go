package main

import (
	"fmt"
	"math"
)

func main() {
	var x, epsilon float64
	fmt.Print("X = ")
	fmt.Scan(&x)
	for i := 0; i < 6; i++ {
		fmt.Scan(&epsilon)
		fmt.Printf("%.7f\n\n", ln1(x, epsilon))
	}
}

func ln1(x, e float64) float64 {
	result := 0.0
	element := x
	trigger := 1.0
	for i := 2; math.Abs(element) > e; i++ {
		result += trigger * element
		trigger *= -1
		element = power2(x, i) / float64(i)
	}
	return result
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