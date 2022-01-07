package main

import "fmt"

func main() {
	var n, a, sum_between_two_zeros, temp_sum int
	fmt.Print("N = ")
	fmt.Scan(&n)
	can_add := false
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if can_add {
			sum_between_two_zeros += a
		}
		if a == 0 {
			if !can_add {
				can_add = true
			} else {
				temp_sum = sum_between_two_zeros
				sum_between_two_zeros = 0
			}
		}
	}
	fmt.Printf("\n%d\n", temp_sum)
}