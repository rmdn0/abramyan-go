package main

import "fmt"

func main() {
	var percent float64
	fmt.Print("P = ")
	fmt.Scan(&percent)
	run_distance := 10.0
	total_distance := 10.0
	days_count := 1
	for total_distance <= 200 {
		run_distance += run_distance * percent / 100
		total_distance += run_distance
		days_count++
	}
	fmt.Printf("\nK = %d\nS = %.3f\n", days_count, total_distance)
}