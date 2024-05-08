package main

import "fmt"

func main() {
	var a int
	fmt.Print("A = ")
	fmt.Scan(&a)
	s := ""
	
	if a % 2 == 0 {
		s += "четное"
	} else {
		s += "нечетное"
	}

	if a > 0 && a < 10 {
		s += " однозначное"
	} else if a > 9 && a < 100 {
		s += " двузначное"
	}  else if a > 99 && a < 1000 {
		s += " трехзначное"
	}
	s += " число"
	fmt.Printf("\n%s\n", s)
}