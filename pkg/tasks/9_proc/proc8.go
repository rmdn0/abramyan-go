package main

import "fmt"

func main() {
	var k, d int
	fmt.Print("K = ")
	fmt.Scan(&k)
	for i := 0; i < 2; i++ {
		fmt.Print("D = ")
		fmt.Scan(&d)
		addRightDigit(d, &k)
		fmt.Printf("K = %d\n", k)
	}
}

func addRightDigit(d int, k *int) {
	*k *= 10
	*k += d
}