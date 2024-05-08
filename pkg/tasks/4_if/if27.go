package main

import "fmt"

func main() {
	var x float32
	fmt.Print("x = ")
	fmt.Scan(&x)
	fmt.Printf("\nf(x) = %d\n", f(x))
}

func f(x float32) int {
	if x < 0 {
		return 0
	} else if int(x) % 2 == 0 {
		return 1
	} else {
		return -1
	}
}