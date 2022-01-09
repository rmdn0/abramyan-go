package main

import "fmt"

func main() {
	var k int
	fmt.Print("K = ")
	fmt.Scan(&k)
	var prev_a, next_a, sets_count int
	var is_increasing bool
	for i := 0; i < k; i++ {
		is_increasing = true
		fmt.Scan(&prev_a)
		fmt.Scan(&next_a)
		for next_a != 0 {
			if is_increasing {
				if prev_a >= next_a {
					is_increasing = false
				}
				prev_a = next_a
			}
			fmt.Scan(&next_a)
		}
		if is_increasing {
			sets_count++
		}
	}
	fmt.Printf("Number of sets with increasing elements: %d\n", sets_count)
}