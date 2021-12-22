package main

import "fmt"

func main() {
	var sweets_cost float64
	fmt.Print("Sweets cost: ")
	fmt.Scan(&sweets_cost)
	fmt.Println()
	for i := 12.0; i <= 20; i += 2 {
		fmt.Printf("%.2f\t", sweets_cost * i / 10)
	}
	fmt.Println()
}