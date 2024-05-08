package main

import (
	"fmt"
	"math"
)

func main() {
	var p, a, b, c float64
	fmt.Print("P = ")
	fmt.Scan(&p)
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	fmt.Printf("\n%.2f\n", power3(a, p))
	fmt.Printf("%.2f\n", power3(b, p))
	fmt.Printf("%.2f\n", power3(c, p))
}

func power1(a, b float64) float64 {
	if a <= 0 {
		return 0
	}
	return math.Exp(b * math.Log(a))
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

func power3(a, b float64) float64 {
	if b - float64(int(b)) == 0 {
		return power2(a, int(b))
	} else {
		return power1(a, b)
	}
}