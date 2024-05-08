package main

import "fmt"

func main() {
	var y int
	for i := 0; i < 5; i++ {
		fmt.Print("Y = ")
		fmt.Scan(&y)
		fmt.Printf("is leap: %t\n", isLeapYear(y))
	}
}

func isLeapYear(y int) bool {
	return y % 4 == 0 && !(y % 100 == 0 && y % 400 != 0)
}