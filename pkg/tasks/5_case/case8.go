package main

import "fmt"

func main() {
	var d, m int
	fmt.Print("D = ")
	fmt.Scan(&d)
	fmt.Print("M = ")
	fmt.Scan(&m)
	d--
	if d == 0 {
		m--
		if m == 0 {
			m = 12
		}
		switch m {
		case 1, 3, 5, 7, 8, 10, 12:
			d = 31
		case 2:
			d = 28
		case 4, 6, 9, 11:
			d = 30
		}
	}
	fmt.Printf("\nD = %d\nM = %d\n", d, m)
}