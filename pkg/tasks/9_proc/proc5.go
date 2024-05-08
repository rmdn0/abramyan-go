package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2, p, s float64
	for i := 0; i < 3; i++ {
		fmt.Scan(&x1, &y1, &x2, &y2)
		rectPS(x1, y1, x2, y2, &p, &s)
		fmt.Printf("%.2f %.2f\n", p, s)
	}
}

func rectPS(x1, y1, x2, y2 float64, p, s *float64) {
	a := math.Abs(x1 - x2)
	b := math.Abs(y1 - y2)
	*p = 2 * (a + b)
	*s = a * b
}