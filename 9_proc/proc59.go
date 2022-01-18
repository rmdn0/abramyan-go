package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3, x4, y4 float64
	fmt.Print("P: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("A: ")
	fmt.Scan(&x2, &y2)
	fmt.Print("B: ")
	fmt.Scan(&x3, &y3)
	fmt.Print("C: ")
	fmt.Scan(&x4, &y4)
	fmt.Printf("\nD(p,ab) = %.2f\n", dist(x1, y1, x2, y2, x3, y3))
	fmt.Printf("D(p,ac) = %.2f\n", dist(x1, y1, x2, y2, x4, y4))
	fmt.Printf("D(p,bc) = %.2f\n", dist(x1, y1, x3, y3, x4, y4))
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