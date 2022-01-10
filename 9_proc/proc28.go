package main

import "fmt"

func main() {
	var k, prime_numbers_counter int
	for i := 0; i < 10; i++ {
		fmt.Scan(&k)
		if isPrime(k) {
			prime_numbers_counter++
		}
	}
	fmt.Printf("\nPrime numbers count: %d\n", prime_numbers_counter)
}

func isPrime(k int) bool {
	divider := 2
	is_prime := true
	for divider < k {
		if k % divider == 0 {
			is_prime = false
			break
		}
		divider++
	}
	return is_prime
}
