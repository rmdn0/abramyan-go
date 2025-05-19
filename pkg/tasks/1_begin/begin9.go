package begin

import (
	"fmt"
	"math"
)

func Begin9() {
	var a, b float32
	fmt.Println("Введите два неотрицательных числа:")
	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	fmt.Println("Среднее геометрическое введенных чисел:")
	fmt.Printf("%.2f\n", math.Sqrt(float64(a*b)))
}
