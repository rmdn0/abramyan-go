package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	temp := 0
	f1 := 1
	f2 := 1
	n_is_fibonacci := false
	for f2 <= n {
		if f2 == n {
			n_is_fibonacci = true
			break
		}
		temp = f1 + f2
		f1 = f2
		f2 = temp
	}
	fmt.Printf("\n%t\n", n_is_fibonacci)
}