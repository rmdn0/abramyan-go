package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Print("A: ")
	fmt.Scan(&x1, &y1)
	for i := 0; i < 3; i++ {
		fmt.Print("next: ")
		fmt.Scan(&x2, &y2)
		fmt.Printf("|A,next| = %.2f\n\n", leng(x1, y1, x2, y2))
	}
}

func leng(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1 - x2, 2) + math.Pow(y1 - y2, 2))
}