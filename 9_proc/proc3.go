package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d, aMean, gMean float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)
	mean(a, b, &aMean, &gMean)
	fmt.Printf("%.2f %.2f\n", aMean, gMean)
	mean(a, c, &aMean, &gMean)
	fmt.Printf("%.2f %.2f\n", aMean, gMean)
	mean(a, d, &aMean, &gMean)
	fmt.Printf("%.2f %.2f\n", aMean, gMean)
}

func mean(x float64, y float64, aMean *float64, gMean *float64) {
	*aMean = (x + y) / 2
	*gMean = math.Sqrt(x * y)
}