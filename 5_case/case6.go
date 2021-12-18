package main

import "fmt"

func main() {
	var length_unit int
	fmt.Print("Length unit no.: ")
	fmt.Scan(&length_unit)
	var length float32
	fmt.Print("Length: ")
	fmt.Scan(&length)
	var length_in_meters float32
	switch length_unit {
	case 1:
		length_in_meters = length / 10
	case 2:
		length_in_meters = length * 1000
	case 3:
		length_in_meters = length
	case 4:
		length_in_meters = length / 1000
	case 5:
		length_in_meters = length / 100
	}
	fmt.Printf("\nLength in meters: %f\n", length_in_meters)
}