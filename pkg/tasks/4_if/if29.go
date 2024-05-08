package main

import "fmt"

func main() {
	var a int
	fmt.Print("A = ")
	fmt.Scan(&a)
	s := ""
	if a == 0 {
		s += "нулевое"
	} else {
		if a > 0 {
			s += "положительное"
		} else {
			s += "отрицательное"
		}

		if a % 2 == 0 {
			s += " четное"
		} else {
			s += " нечетное"
		}
	}
	s += " число"
	fmt.Printf("\n%s\n", s)
}