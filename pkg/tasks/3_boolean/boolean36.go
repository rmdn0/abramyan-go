package main

import "fmt"

func main() {
	var x1, y1, x2, y2 int
	fmt.Print("x1, y1 = ")
	fmt.Scan(&x1, &y1)
	fmt.Print("x2, y2 = ")
	fmt.Scan(&x2, &y2)
	fmt.Print("\nЛадья может за один ход перейти с одного поля на другое: ")
	fmt.Printf("%t\n", x1 == x2 || y1 == y2)
}