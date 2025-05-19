package begin

import "fmt"

func Begin3() {
	var a, b float32
	fmt.Print("Введите стороны прямоугольника:\na = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	s := a * b
	p := 2 * (a + b)
	fmt.Println(
		"Площадь и периметр прямоугольника со сторонами a и b:\nS =",
		fmt.Sprintf("%.2f", s),
		"\nP =",
		fmt.Sprintf("%.2f", p),
	)
}
