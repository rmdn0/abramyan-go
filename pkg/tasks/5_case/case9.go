package main

import "fmt"

func main() {
		var d, m int
	fmt.Print("D = ")
	fmt.Scan(&d)
	fmt.Print("M = ")
	fmt.Scan(&m)
	d++
	switch m {
	case 1, 3, 5, 7, 8, 10, 12:
		if d == 32 {
			d = 1
			if m == 12 {
				m = 1
			} else {
				m++
			}
		}
	case 2:
		if d == 29 {
			d = 1
			m++
		}
	case 4, 6, 9, 11:
		if d == 31 {
			d = 1
			m++
		}
	}
	fmt.Printf("\nD = %d\nM = %d\n", d, m)
}