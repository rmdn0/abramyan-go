package main

import "fmt"

func main() {
	var day, month int
	fmt.Print("Day: ")
	fmt.Scan(&day)
	fmt.Print("Month: ")
	fmt.Scan(&month)
	var zodiak_sign string
	switch month {
	case 1:
		switch {
		case day <= 19:
			zodiak_sign = "Козерог"
		default:
			zodiak_sign = "Водолей"
		}
	case 2:
		switch {
		case day <= 18:
			zodiak_sign = "Водолей"
		default:
			zodiak_sign = "Рыбы"
		}
	case 3:
		switch {
		case day <= 20:
			zodiak_sign = "Рыбы"
		default:
			zodiak_sign = "Овен"
		}
	case 4:
		switch {
		case day <= 19:
			zodiak_sign = "Овен"
		default:
			zodiak_sign = "Телец"
		}
	case 5:
		if day <= 20 {
			zodiak_sign = "Телец"
		} else {
			zodiak_sign = "Близнецы"
		}
	case 6:
		if day <= 21 {
			zodiak_sign = "Близнецы"
		} else {
			zodiak_sign = "Рак"
		}
	case 7:
		if day <= 22 {
			zodiak_sign = "Рак"
		} else {
			zodiak_sign = "Лев"
		}
	case 8:
		if day <= 22 {
			zodiak_sign = "Лев"
		} else {
			zodiak_sign = "Дева"
		}
	case 9:
		switch {
		case day <= 22:
			zodiak_sign = "Дева"
		default:
			zodiak_sign = "Весы"
		}
	case 10:
		switch {
		case day <= 22:
			zodiak_sign = "Весы"
		default:
			zodiak_sign = "Скорпион"
		}
	case 11:
		switch {
		case day <= 22:
			zodiak_sign = "Скорпион"
		default:
			zodiak_sign = "Стрелец"
		}
	case 12:
		switch {
		case day <= 21:
			zodiak_sign = "Стрелец"
		default:
			zodiak_sign = "Козерог"
		}
	}
	fmt.Printf("\n%s\n", zodiak_sign)
}