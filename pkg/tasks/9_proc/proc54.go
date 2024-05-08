package main

import "fmt"

func main() {
	var d, m, y int	
	for i := 0; i < 3; i++ {
		fmt.Print("D = ")
		fmt.Scan(&d)
		fmt.Print("M = ")
		fmt.Scan(&m)
		fmt.Print("Y = ")
		fmt.Scan(&y)
		prevDate(&d, &m, &y)
		fmt.Printf("prev date: d = %d m = %d y = %d\n", d, m, y)
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

func prevDate(d, m, y *int) {
	*d--
	if *d == 0 {
		*m--
		if *m == 0 {
			*y--
			*m = 12
		}
		*d = monthDays(*m, *y)
	}
}