package main

import "fmt"

func main() {
	var number int
	fmt.Print("Number (100-999): ")
	fmt.Scan(&number)
	var string_number string
	hundreds := number / 100
	dozens := number % 100 / 10
	units := number % 10

	switch hundreds {
	case 1:
		string_number = "сто"
	case 2:
		string_number = "двести"
	case 3:
		string_number = "триста"
	case 4:
		string_number = "четыреста"
	case 5:
		string_number = "пятьсот"
	case 6:
		string_number = "шестьсот"
	case 7:
		string_number = "семьсот"
	case 8:
		string_number = "восемьсот"
	case 9:
		string_number = "девятьсот"
	}

	switch dozens {
	case 1:
		switch units {
		case 0:
			string_number += " десять"
		case 1:
			string_number += " одиннадцать"
		case 2:
			string_number += " двенадцать"
		case 3:
			string_number += " тринадцать"
		case 4:
			string_number += " четырнадцать"
		case 5:
			string_number += " пятнадцать"
		case 6:
			string_number += " шестнадцать"
		case 7:
			string_number += " семнадцать"
		case 8:
			string_number += " восемнадцать"
		case 9:
			string_number += " девятнадцать"
		}
	case 2:
		string_number += " двадцать"
	case 3:
		string_number += " тридцать"
	case 4:
		string_number += " сорок"
	case 5:
		string_number += " пятьдесят"
	case 6:
		string_number += " шестьдесят"
	case 7:
		string_number += " семьдесят"
	case 8:
		string_number += " восемьдесят"
	case 9:
		string_number += " девяносто"
	}

	if dozens != 1 {
		switch units {
		case 1:
			string_number += " один"
		case 2:
			string_number += " два"
		case 3:
			string_number += " три"
		case 4:
			string_number += " четыре"
		case 5:
			string_number += " пять"
		case 6:
			string_number += " шесть"
		case 7:
			string_number += " семь"
		case 8:
			string_number += " восемь"
		case 9:
			string_number += " девять"
		}
	}
	fmt.Printf("\n%s\n", string_number)
}