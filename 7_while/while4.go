package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	// 1st way
	// for n >= 3 {
	// 	if (n%3==0) { 
	// 		n /= 3
	// 	} else {
	// 		break
	// 	}
	// }
	// fmt.Printf("\n%t\n", n == 1)
	// 2nd way
	power := 1
	for power < n {
		power *= 3
	}
	fmt.Printf("\n%t\n", power == n)
}