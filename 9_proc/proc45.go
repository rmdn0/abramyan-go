package main

import (
	"fmt"
	"math"
)

func main() {
	var x, a, epsilon float64
	fmt.Print("x = ")
	fmt.Scan(&x)
	fmt.Print("a = ")
	fmt.Scan(&a)
	for i := 0; i < 6; i++ {
		fmt.Print("e = ")
		fmt.Scan(&epsilon)
		fmt.Printf("%.7f\n\n", power4(x, a, epsilon))
	}
}

func power4(x, a, epsilon float64) float64 {
	result := 0.0
	element := 1 + a * x
	for i := 2; math.Abs(element) > epsilon; i++ {
		result += element
		element = a * power2(x, i) / fact(i)
		for j := i - 1; j > 0; j-- {
			element *= a - float64(j)
		}
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