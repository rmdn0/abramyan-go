package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var prev_a, curr_a, next_a float64
	fmt.Scan(&prev_a)
	fmt.Scan(&curr_a)
	sawtooth_sequence_intruder_index := 0
	for i := 2; i < n; i++ {
		fmt.Scan(&next_a)
		if !(curr_a > prev_a && curr_a > next_a || curr_a < prev_a && curr_a < next_a) && sawtooth_sequence_intruder_index == 0 {
			sawtooth_sequence_intruder_index = i
		}
		prev_a = curr_a
		curr_a = next_a
	}
	fmt.Printf("\n%d\n", sawtooth_sequence_intruder_index)
}