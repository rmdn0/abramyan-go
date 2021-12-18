package main

import "fmt"

func main() {
	var mark int
	fmt.Print("Mark: ")
	fmt.Scan(&mark)
	var mark_in_str string
	switch mark {
	case 1:
		mark_in_str = "плохо"
	case 2:
		mark_in_str = "неудовлитворительно"
	case 3:
		mark_in_str = "удовлитворительно"
	case 4:
		mark_in_str = "хорошо"
	case 5:
		mark_in_str = "отлично"
	default:
		mark_in_str = "ошибка"
	}
	fmt.Printf("\n%s\n", mark_in_str)
}