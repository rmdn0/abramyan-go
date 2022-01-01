package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	k := 0
	for a != 0 {
		fmt.Scan(&a)
		k++
	}
	fmt.Printf("\n%d\n", k)
}