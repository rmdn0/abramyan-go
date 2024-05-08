package main

import "fmt"

func main() {
	var k int
	fmt.Print("K = ")
	fmt.Scan(&k)
	var prev_a, curr_a, next_a, sawtooth_sequence_intruder_index, j int
	for i := 0; i < k; i++ {
		fmt.Scan(&prev_a)
		fmt.Scan(&curr_a)
		fmt.Scan(&next_a)
		j = 1
		sawtooth_sequence_intruder_index = 0
		for next_a != 0 {
			if sawtooth_sequence_intruder_index == 0 {
				if !(curr_a > prev_a && curr_a > next_a || curr_a < prev_a && curr_a < next_a) {
					sawtooth_sequence_intruder_index = j + 1
				}
				prev_a = curr_a
				curr_a = next_a
			}
			j++
			fmt.Scan(&next_a)
		}
		if sawtooth_sequence_intruder_index == 0 {
			fmt.Println(j + 1)
		} else {
			fmt.Println(sawtooth_sequence_intruder_index)
		}
	}
}