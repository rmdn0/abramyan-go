package main

import "fmt"

func main() {
	var a, b, summ int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	for i := a; i <= b; i++ {
		summ += i
	}
	fmt.Printf("\n%d\n", summ)
}