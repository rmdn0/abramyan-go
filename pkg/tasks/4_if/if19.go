package main

import "fmt"

func main() {
	var a, b, c, d float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	fmt.Print("D = ")
	fmt.Scan(&d)
	var index int
	// // 1st way
	// if a == b && b == c {
	// 	index = 4
	// } else if a == b && b == d {
	// 	index = 3
	// } else if a == c && c == d {
	// 	index = 2
	// } else if b == c && c == d {
	// 	index = 1
	// }
	// 2nd way
	if a == b {
		if b == c {
			index = 4
		} else {
			index = 3
		}
	} else if b == c {
		index = 1
	} else if a == c {
		index = 2
	}
	fmt.Printf("\nThe index of different number: %d\n", index)
}