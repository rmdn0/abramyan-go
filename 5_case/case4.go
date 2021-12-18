package main

import "fmt"

func main() {
	var month_number int
	fmt.Print("Month no: ")
	fmt.Scan(&month_number)
	var days_count int
	switch month_number {
	case 1, 3, 5, 7, 8, 10, 12:
		days_count = 31
	case 2:
		days_count = 28
	case 4, 6, 9, 11:
		days_count = 30
	}
	fmt.Printf("\nDays count %d\n", days_count)
}