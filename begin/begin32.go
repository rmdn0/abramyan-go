package main

import "fmt"

func main() {
	var tf, tc float32
	fmt.Print("T(c) = ")
	fmt.Scan(&tc)
	tf = tc * 9 / 5 + 32
	fmt.Printf("\nT(f) = %.2f\n", tf)
}