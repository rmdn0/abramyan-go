package main

import "fmt"

func main() {
	var ar [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&ar[i])
	}
	a := 0
	for i := 1; i < 9; i++ {
		if ar[0] < ar[i] && ar[i] < ar[9] {
			a = i + 1
		}
	}
	fmt.Println(a)
}
