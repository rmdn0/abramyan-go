package main

import "fmt"

func main() {
	var a, b, numbers_count int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Println()
	for i := a; i <= b; i++ {
		fmt.Printf("%d\t", i)
		numbers_count++
	}
	fmt.Printf("\nN = %d\n", numbers_count)
}