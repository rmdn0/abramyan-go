package main

import "fmt"

func main() {
	var x1, y1, x2, y2 int
	fmt.Print("x1, y1 = ")
	fmt.Scan(&x1, &y1)
	fmt.Print("x2, y2 = ")
	fmt.Scan(&x2, &y2)
	fmt.Print("\nBoth of fields are the same color: ")
	fmt.Printf("%t\n", both_are_black(x1, y1, x2, y2) || both_are_white(x1, y1, x2, y2))
}

func is_odd(x int) bool {
	return x % 2 == 1
}

func both_are_black(x1, y1, x2, y2 int) bool {
	return is_black(x1, y1) && is_black(x2, y2)
}

func is_black(x, y int) bool {
	return is_odd(x) && is_odd(y) || !is_odd(x) && !is_odd(y)
}

func both_are_white(x1, y1, x2, y2 int) bool {
	return is_white(x1, y1) && is_white(x2, y2)
}

func is_white(x, y int) bool {
	return !is_odd(x) && is_odd(y) || is_odd(x) && !is_odd(y)
}