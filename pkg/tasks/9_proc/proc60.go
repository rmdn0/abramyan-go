package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3, x4, y4, h1, h2, h3 float64
	fmt.Print("A: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("B: ")
	fmt.Scan(&x2, &y2)
	fmt.Print("C: ")
	fmt.Scan(&x3, &y3)
	fmt.Print("D: ")
	fmt.Scan(&x4, &y4)
	altitudes(x1, y1, x2, y2, x3, y3, &h1, &h2, &h3)
	fmt.Printf("\nABC: ha = %.2f hb = %.2f hc = %.2f\n", h1, h2, h3)
	altitudes(x1, y1, x2, y2, x4, y4, &h1, &h2, &h3)
	fmt.Printf("ABD: ha = %.2f hb = %.2f hc = %.2f\n", h1, h2, h3)
	altitudes(x1, y1, x3, y3, x4, y4, &h1, &h2, &h3)
	fmt.Printf("ACD: ha = %.2f hb = %.2f hc = %.2f\n", h1, h2, h3)
}

func leng(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1 - x2, 2) + math.Pow(y1 - y2, 2))
}

func perim(x1, y1, x2, y2, x3, y3 float64) float64 {
	return leng(x1, y1, x2, y2) + leng(x2, y2, x3, y3) + leng(x3, y3, x1, y1)
}

func area(x1, y1, x2, y2, x3, y3 float64) float64 {
	p := perim(x1, y1, x2, y2, x3, y3) / 2
	a := leng(x1, y1, x2, y2)
	b := leng(x2, y2, x3, y3)
	c := leng(x3, y3, x1, y1)
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func dist(x1, y1, x2, y2, x3, y3 float64) float64 {
	return 2 * area(x1, y1, x2, y2, x3, y3) / leng(x2, y2, x3, y3)
}

func altitudes(x1, y1, x2, y2, x3, y3 float64, h1, h2, h3 *float64) {
	*h1 = dist(x1, y1, x2, y2, x3, y3)
	*h2 = dist(x2, y2, x1, y1, x3, y3)
	*h3 = dist(x3, y3, x1, y1, x2, y2)
}