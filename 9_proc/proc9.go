package main

import "fmt"

func main() {
	var k, d int
	fmt.Print("K = ")
	fmt.Scan(&k)
	for i := 0; i < 2; i++ {
		fmt.Print("D = ")
		fmt.Scan(&d)
		addLeftDigit(d, &k)
		fmt.Printf("K = %d\n", k)
	}
}

func addLeftDigit(d int, k *int) {
	invDigits(k)	
	for *k > 0 {
		d *= 10
		d += *k % 10
		*k /= 10
	}
	*k = d
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