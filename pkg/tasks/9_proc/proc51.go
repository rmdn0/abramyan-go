package main

import "fmt"

func main() {
	var h, m, s, t int
	fmt.Print("H = ")
	fmt.Scan(&h)
	fmt.Print("M = ")
	fmt.Scan(&m)
	fmt.Print("S = ")
	fmt.Scan(&s)
	fmt.Print("T = ")
	fmt.Scan(&t)
	incTime(&h, &m, &s, t)
	fmt.Printf("\nH = %d\nM = %d\nS = %d\n", h, m, s)
}

func incTime(h, m, s *int, t int) {
	*s += t
	*m += *s / 60
	*s %= 60
	*h += *m / 60
	*m %= 60
}