package main

import "fmt"

func main() {
	var direction rune
	fmt.Print("Direction: ")
	fmt.Scanf("%c", &direction)
	var command1, command2 int
	fmt.Print("N1 = ")
	fmt.Scan(&command1)
	fmt.Print("N2 = ")
	fmt.Scan(&command2)
	switch command1 {
	case -1:
		switch direction {
		case 'С':
			direction = 'В'
		case 'В':
			direction = 'Ю'
		case 'Ю':
			direction = 'З'
		case 'З':
			direction = 'С'
		}
	case 1:
		switch direction {
		case 'С':
			direction = 'З'
		case 'З':
			direction = 'Ю'
		case 'Ю':
			direction = 'В'
		case 'В':
			direction = 'С'
		}
	case 2:
		switch direction {
		case 'С':
			direction = 'Ю'
		case 'В':
			direction = 'З'
		case 'Ю':
			direction = 'С'
		case 'З':
			direction = 'В'
		}
	}
	switch command2 {
	case -1:
		switch direction {
		case 'С':
			direction = 'В'
		case 'В':
			direction = 'Ю'
		case 'Ю':
			direction = 'З'
		case 'З':
			direction = 'С'
		}
	case 1:
		switch direction {
		case 'С':
			direction = 'З'
		case 'З':
			direction = 'Ю'
		case 'Ю':
			direction = 'В'
		case 'В':
			direction = 'С'
		}
	case 2:
		switch direction {
		case 'С':
			direction = 'Ю'
		case 'В':
			direction = 'З'
		case 'Ю':
			direction = 'С'
		case 'З':
			direction = 'В'
		}
	}
	fmt.Printf("\nNew direction: %c\n", direction)
}