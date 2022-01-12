package main

import "fmt"
const pi = 3.14

func main() {
	var d float64
	for i := 0; i < 5; i++ {
		fmt.Print("D = ")
		fmt.Scan(&d)
		fmt.Printf("%.2f\n\n", degToRad(d))
	}
}

func degToRad(d float64) float64 {
	return d * pi / 180
}