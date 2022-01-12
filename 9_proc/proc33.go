package main

import "fmt"
const pi = 3.14

func main() {
	var r float64
	for i := 0; i < 5; i++ {
		fmt.Print("R = ")
		fmt.Scan(&r)
		fmt.Printf("%.2f\n\n", radToDeg(r))
	}
}

func radToDeg(r float64) float64 {
	return r * 180 / pi
}