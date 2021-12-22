package main

import "fmt"

func main() {
	var numeric_excercises_count int
	fmt.Print("Excercises count (numeric): ")
	fmt.Scan(&numeric_excercises_count)
	var string_excercises_count string
	unit := numeric_excercises_count % 10
	switch numeric_excercises_count / 10 {
	case 1:
		switch unit {
		case 0:
			string_excercises_count = "десять"
		case 1:
			string_excercises_count = "одиннадцать"
		case 2:
			string_excercises_count = "двенадцать"
		case 3:
			string_excercises_count = "тринадцать"
		case 4:
			string_excercises_count = "четырнадцать"
		case 5:
			string_excercises_count = "пятнадцать"
		case 6:
			string_excercises_count = "шестнадцать"
		case 7:
			string_excercises_count = "семнадцать"
		case 8:
			string_excercises_count = "восемьнадцать"
		case 9:
			string_excercises_count = "девятнадцать"
		}
		string_excercises_count += " учебных заданий"
	case 2:
		string_excercises_count = "двадцать"
	case 3:
		string_excercises_count = "тридцать"
	case 4:
		string_excercises_count = "сорок"
	}
	if numeric_excercises_count >= 20 {
		switch {
		case unit == 1:
			string_excercises_count += " одно"
			string_excercises_count += " учебное задание"
		case unit > 1 && unit < 5:
			switch unit {
			case 2:
				string_excercises_count += " два"
			case 3:
				string_excercises_count += " три"
			case 4:
				string_excercises_count += " четыре"
			}
			string_excercises_count += " учебных задания"
		default:
			switch unit {
			case 5:
				string_excercises_count += " пять"
			case 6:
				string_excercises_count += " шесть"
			case 7:
				string_excercises_count += " семь"
			case 8:
				string_excercises_count += " восемь"
			case 9:
				string_excercises_count += " девять"
			}
			string_excercises_count += " учебных заданий"
		}
	}
	fmt.Printf("\n%s\n", string_excercises_count)
}