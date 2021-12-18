package main

import "fmt"

func main() {
	var weekday_number int
	fmt.Print("Weekday number = ")
	fmt.Scan(&weekday_number)
	day_in_string := ""
	switch weekday_number {
	case 1:
		day_in_string = "понедельник"
	case 2:
		day_in_string = "вторник"
	case 3:
		day_in_string = "среда"
	case 4:
		day_in_string = "четверг"
	case 5:
		day_in_string = "пятница"
	case 6:
		day_in_string = "суббота"
	case 7:
		day_in_string = "воскресенье"
	}
	fmt.Printf("\n%s\n", day_in_string)
}