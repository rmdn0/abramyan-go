package main

import "fmt"

func main() {
	var n, k, power_n_numbers_counter int
	fmt.Print("N = ")
	fmt.Scan(&n)
	for i := 0; i < 10; i++ {
		fmt.Scan(&k)
		if isPowerN(k, n) {
			power_n_numbers_counter++
		}
	}
	fmt.Printf("\nPower of N numbers count: %d\n", power_n_numbers_counter)
}

func isPowerN(k, n int) bool {
	power := 1
	for power < k {
		power *= n
	}
	return power == k
}
