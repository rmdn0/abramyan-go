package main

import "fmt"

func main() {
	var year_no int
	fmt.Print("Year no: ")
	fmt.Scan(&year_no)
	fmt.Printf("\nDays in this year: %d\n", number_of_days_in_a_year(year_no))
}

func number_of_days_in_a_year(x int) int {
	// 1st way
	if x % 4 == 0 {
		if x % 100 == 0 && x % 400 != 0 {
			return 365
		}
		return 366
	}
	return 365

	// 2nd way
	// if x % 4 == 0 && x % 100 != 0 && x % 400 != 0 {
	// 	return 366
	// }
	// return 365
}