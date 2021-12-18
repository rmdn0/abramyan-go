package main

import "fmt"

func main() {
	var month_number int
	fmt.Print("Month no: ")
	fmt.Scan(&month_number)
	season := ""
	switch {
	case month_number > 0 && month_number < 3 || month_number == 12:
		season = "winter"
	case month_number > 2 && month_number < 6:
		season = "spring"
	case month_number > 5 && month_number < 9:
		season = "summer"
	case month_number > 8 && month_number < 12:
		season = "autumn"
	}
	fmt.Printf("\nIt's %s\n", season)
}