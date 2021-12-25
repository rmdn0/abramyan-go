package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	f1 := 1
	f2 := 1
	next_f := 0
	fmt.Println()
	fmt.Printf("%d\t%d\t", f1, f2)
	for i := 2; i < n; i++ {
		next_f = f1 + f2
		fmt.Printf("%d\t", next_f)
		f1 = f2
		f2 = next_f
	}
	fmt.Println()
}