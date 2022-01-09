package main

import "fmt"

func main() {
	var k int
	fmt.Print("K = ")
	fmt.Scan(&k)
	var prev_a, next_a, sets_count int
	var is_increasing, is_decreasing bool
	for i := 0; i < k; i++ {
		is_increasing = true
		is_decreasing = true
		fmt.Scan(&prev_a)
		fmt.Scan(&next_a)
		for next_a != 0 {
			if is_increasing || is_decreasing {
				if prev_a >= next_a {
					is_increasing = false
				}
				if prev_a <= next_a {
					is_decreasing = false
				}
				prev_a = next_a
			}
			fmt.Scan(&next_a)
		}
		if is_increasing || is_decreasing {
			sets_count++
		}
	}
	fmt.Printf("Number of sets with increasing or decreasing elements: %d\n", sets_count)
}