package main

import "fmt"
import "math"

func main() {
	var x1, y1, x2, y2 int
	fmt.Print("x1, y1 = ")
	fmt.Scan(&x1, &y1)
	fmt.Print("x2, y2 = ")
	fmt.Scan(&x2, &y2)
	elephant := math.Abs(float64(x1 - x2)) == math.Abs(float64(y1 - y2))
	king := int(math.Abs(float64(x1 - x2))) <= 1 && int(math.Abs(float64(y1 - y2))) <= 1
	ladya := x1 == x2 || y1 == y2
	fmt.Print("\nФерзь может за один ход перейти с одного поля на другое: ")
	fmt.Printf("%t\n", elephant || king || ladya)
}