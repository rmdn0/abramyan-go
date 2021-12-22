package main

import "fmt"

func main() {
	var sweets_cost float32
	fmt.Print("Sweets cost: ")
	fmt.Scan(&sweets_cost)
	fmt.Println()
	for i := 1; i < 11; i++ {
		fmt.Printf("%.2f\t", sweets_cost * float32(i))
	}
	fmt.Println()
}