package main

import "fmt"

func main() {
	var year int
	fmt.Print("Year: ")
	fmt.Scan(&year)
	year = (year - 4) % 60
	group_number := year / 12 + 1
	animal_number := year % 12 + 1

	var animal string
	switch animal_number {
	case 1:
		animal = " крысы"
	case 2:
		animal = " коровы"
	case 3:
		animal = " тигра"
	case 4:
		animal = " зайца"
	case 5:
		animal = " дракона"
	case 6:
		animal = " змеи"
	case 7:
		animal = " лошади"
	case 8:
		animal = " овцы"
	case 9:
		animal = " обезьяны"
	case 10:
		animal = " курицы"
	case 11:
		animal = " собаки"
	case 12:
		animal = " свиньи"
	}

	var group string
	switch (group_number) {
	case 1:
		switch animal_number {
		case 3, 4, 5 :
			group = " зеленого"
		default:
			group = " зеленой"
		}
	case 2:
		switch animal_number {
		case 3, 4, 5 :
			group = " красного"
		default:
			group = " красной"
		}
	case 3:
		switch animal_number {
		case 3, 4, 5 :
			group = " желтого"
		default:
			group = " желтой"
		}
	case 4:
		switch animal_number {
		case 3, 4, 5 :
			group = " белого"
		default:
			group = " белой"
		}
	case 5:
		switch animal_number {
		case 3, 4, 5 :
			group = " черного"
		default:
			group = " черной"
		}
	}
	string_east_year := "год" + group + animal
	fmt.Printf("\n%s\n", string_east_year)
}