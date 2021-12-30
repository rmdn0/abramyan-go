package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	temp := 0
	f1 := 1
	f2 := 1
	k := 1
	for f2 <= n {
		k++
		temp = f1 + f2
		f1 = f2
		f2 = temp
	}
	fmt.Printf("\n%d\n", k)
}