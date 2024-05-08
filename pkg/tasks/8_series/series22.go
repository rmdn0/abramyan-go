package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a float64
	fmt.Scan(&a)
	prev_a := a
	var decreasing_sequence_intruder_index int
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if prev_a <= a {
			decreasing_sequence_intruder_index = i + 1
		}
		prev_a = a
	}
	fmt.Printf("\n%d\n", decreasing_sequence_intruder_index)
}