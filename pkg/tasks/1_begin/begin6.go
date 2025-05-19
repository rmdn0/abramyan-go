package begin

import "fmt"

func Begin6() {
	var a, b, c float32
	fmt.Println("Введите ребра прямоугольного параллелепипеда:")
	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	fmt.Print("c = ")
	fmt.Scan(&c)

	v := a * b * c
	s := 2 * (a*b + b*c + c*a)
	fmt.Println("\nОбъем и площадь поверхности:")
	fmt.Println("V =", fmt.Sprintf("%.3f", v), "\nS =", fmt.Sprintf("%.3f", s))
}
