package main

import "fmt"

func main() {
	var k int
	for i := 0; i < 5; i++ {
		fmt.Print("K = ")
		fmt.Scan(&k)
		invDigits(&k)
		fmt.Printf("K = %d\n", k)
	}
}

func invDigits(k *int) {
	var new_k int
	for *k > 0 {
		new_k *= 10
		new_k += *k % 10
		*k /= 10
	}
	*k = new_k
}