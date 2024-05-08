package main

import "fmt"

func main() {
	var t1, t2, u, v float32
	fmt.Print("V = ")
	fmt.Scan(&v)
	fmt.Print("U = ")
	fmt.Scan(&u)
	fmt.Print("T1 = ")
	fmt.Scan(&t1)
	fmt.Print("T2 = ")
	fmt.Scan(&t2)
	s := v * t1 + (v - u) * t2
	fmt.Printf("\nS = %.2f\n", s)
}