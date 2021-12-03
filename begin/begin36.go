package main

import "fmt"

func main() {
	var v1, v2, s, t float32
	fmt.Print("V1 = ")
	fmt.Scan(&v1)
	fmt.Print("V2 = ")
	fmt.Scan(&v2)
	fmt.Print("S = ")
	fmt.Scan(&s)
	fmt.Print("T = ")
	fmt.Scan(&t)
	total_s := s + t * (v1 + v2)
	fmt.Printf("\nTotal S = %.2f\n", total_s)
}