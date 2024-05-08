package main

import "fmt"

func main() {
	var numeric_age int
	fmt.Print("Age (numeric): ")
	fmt.Scan(&numeric_age)
	var string_age string
	switch numeric_age / 10 {
	case 2:
		string_age = "двадцать"
	case 3:
		string_age = "тридцать"
	case 4:
		string_age = "сорок"
	case 5:
		string_age = "пятдесят"
	case 6:
		string_age = "шестдесят"
	}
	switch numeric_age % 10 {
	case 0:
		string_age += " лет"
	case 1:
		string_age += " один год"
	case 2:
		string_age += " два года"
	case 3:
		string_age += " три года"
	case 4:
		string_age += " четыре года"
	case 5:
		string_age += " пять лет"
	case 6:
		string_age += " шесть лет"
	case 7:
		string_age += " семь лет"
	case 8:
		string_age += " восемь лет"
	case 9:
		string_age += " девять лет"
	}
	fmt.Printf("\n%s\n", string_age)
}