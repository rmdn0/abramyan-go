package main

import (
	"fmt"
	"math"
)

func main() {
	var a, p, s float64
	for i := 0; i < 3; i++ {
		fmt.Scan(&a)
		trianglePS(a, &p, &s)
		fmt.Printf("%.2f %.2f\n", p, s)
	}
}

func trianglePS(a float64, p *float64, s *float64) {
	*p = 3 * a
	*s = math.Pow(a, 2) * math.Sqrt(3) / 4
}