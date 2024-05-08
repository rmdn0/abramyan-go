package main

import "fmt"

func main() {
	var y, m int
	fmt.Print("Y = ")
	fmt.Scan(&y)	
	for i := 0; i < 3; i++ {
		fmt.Print("M = ")
		fmt.Scan(&m)
		fmt.Printf("month days: %d\n", monthDays(m, y))
	}
}

func isLeapYear(y int) bool {
	return y % 4 == 0 && !(y % 100 == 0 && y % 400 != 0)
}

func monthDays(m, y int) int {
	var days_count int
	switch m {
	case 1, 3, 5, 7, 8, 10, 12:
		days_count = 31
	case 2:
		if isLeapYear(y) {
			days_count = 29
		} else {
			days_count = 28
		}
	case 4, 6, 9, 11:
		days_count = 30
	}
	return days_count
}