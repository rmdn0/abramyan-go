package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	temp := 0
	for b > 0 {
		temp = b
		b = a % b
		a = temp
	}
	fmt.Printf("\n%d\n", a)
}