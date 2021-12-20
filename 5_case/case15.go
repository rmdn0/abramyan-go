package main

import "fmt"

func main() {
	var value, suit int
	fmt.Print("Value, N = ")
	fmt.Scan(&value)
	fmt.Print("Suit, M = ")
	fmt.Scan(&suit)
	var s string
	switch value {
	case 6:
		s = "шестерка"
	case 7:
		s = "семерка"
	case 8:
		s = "восьмерка"
	case 9:
		s = "девятка"
	case 10:
		s = "десятка"
	case 11:
		s = "валет"
	case 12:
		s = "дама"
	case 13:
		s = "король"
	case 14:
		s = "туз"
	}
	switch suit {
	case 1:
		s += " пик"
	case 2:
		s += " треф"
	case 3:
		s += " бубнов"
	case 4:
		s += " червей"
	}
	fmt.Printf("\n%s\n", s)
}