package main

import "fmt"

func main() {
	var k int
	fmt.Print("K = ")
	fmt.Scan(&k)
	var prev_a, curr_a, next_a, sawtooth_sets_counter int
	var is_sawtooth bool
	for i := 0; i < k; i++ {
		is_sawtooth = true
		fmt.Scan(&prev_a)
		fmt.Scan(&curr_a)
		fmt.Scan(&next_a)
		for next_a != 0 {
			if is_sawtooth {
				if !(curr_a > prev_a && curr_a > next_a || curr_a < prev_a && curr_a < next_a) {
					is_sawtooth = false
				}
				prev_a = curr_a
				curr_a = next_a
			}
			fmt.Scan(&next_a)
		}
		if is_sawtooth {
			sawtooth_sets_counter++
		}
	}
	fmt.Printf("\nNumber of sawtooth sets: %d\n", sawtooth_sets_counter)
}