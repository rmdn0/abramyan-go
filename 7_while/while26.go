package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	temp := 0
	f1 := 1
	f2 := 1
	for f2 <= n {
		temp = f1 + f2
		f1 = f2
		f2 = temp
	}
	fmt.Printf("\n%d\t%d\n", f2 - f1, f2)
}