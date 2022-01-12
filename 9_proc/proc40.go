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
		fmt.Printf("%.7f\n\n", exp1(x, epsilon))
	}
}

func exp1(x, e float64) float64 {
	result := 0.0
	element := 1.0
	for i := 1; element > e; i++ {
		result += element
		element = power2(x, i) / fact(i)
	}
	return result
}

func fact(n int) float64 {
	result := 1.0
	for n > 1 {
		result *= float64(n)
		n--
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