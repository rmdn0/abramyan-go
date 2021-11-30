package main

import "fmt"

func main() {
	const pi float32 = 3.14
	var r, l, s float32
	fmt.Println("Введите радиус круга:")
	fmt.Print("R = ")
	fmt.Scan(&r)
	l = 2 * pi * r
	s = pi * r * r
	fmt.Println("Длина окружности и площадь круга:")
	fmt.Println("L =", fmt.Sprintf("%.3f", l))
	fmt.Println("S =", fmt.Sprintf("%.3f", s))
}