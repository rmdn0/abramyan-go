package main

import "fmt"

func main() {
	var sweets_cost float64
	fmt.Print("Sweets cost: ")
	fmt.Scan(&sweets_cost)
	fmt.Println()
	for i := 1.0; i < 11; i++ {
		fmt.Printf("%.2f\t", sweets_cost * i / 10)
	}
	fmt.Println()
}