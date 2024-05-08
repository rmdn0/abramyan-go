package main

import "fmt"

func main() {
	var k, palindrome_numbers_counter int
	for i := 0; i < 10; i++ {
		fmt.Scan(&k)
		if isPalindrome(k) {
			palindrome_numbers_counter++
		}
	}
	fmt.Printf("\nPalindrome numbers count: %d\n", palindrome_numbers_counter)
}

func isPalindrome(k int) bool {
	reversed_k := reverse(k)
	return reversed_k == k
}

func reverse(k int) int {
	reversed_k := 0
	for k > 0 {
		reversed_k *= 10
		reversed_k += k % 10
		k /= 10
	}
	return reversed_k
}
