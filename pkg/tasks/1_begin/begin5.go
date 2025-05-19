package begin

import "fmt"
import "math"

func Begin5() {
	var a float32
	fmt.Print("Введите ребро куба:\na = ")
	fmt.Scan(&a)
	v := math.Pow(float64(a), 3)
	s := 6 * math.Pow(float64(a), 2)
	fmt.Println("\nV =", fmt.Sprintf("%.3f", v))
	fmt.Println("S =", fmt.Sprintf("%.3f", s))
}
