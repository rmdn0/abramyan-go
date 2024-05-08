package main

import "fmt"

func main() {
	var a, b, c, d, e, f, g, h int
	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	fmt.Print("c = ")
	fmt.Scan(&c)
	fmt.Print("d = ")
	fmt.Scan(&d)
	fmt.Print("e = ")
	fmt.Scan(&e)
	fmt.Print("f = ")
	fmt.Scan(&f)
	fmt.Print("g = ")
	fmt.Scan(&g)
	fmt.Print("h = ")
	fmt.Scan(&h)
	fmt.Println(fracSum(a, b, c, d))
	fmt.Println(fracSum(a, b, e, f))
	fmt.Println(fracSum(a, b, g, h))
}

func gcd2(a, b int) int {
	var temp_a int
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}
	for b != 0 {
		temp_a = a
		a = b
		b = temp_a % b
	}
	return a
}

func frac1(a, b int, p, q *int) {
	if b < 0 {
		a *= -1
		b *= -1
	}
	gcd := gcd2(a, b)
	*p = a / gcd
	*q = b / gcd
}

func fracSum(n1, d1, n2, d2 int) (int, int) {
	frac1(n1, d1, &n1, &d1)
	frac1(n2, d2, &n2, &d2)
	gcd := gcd2(d1, d2) 
	multiplier1 := d2 / gcd
	multiplier2 := d1 / gcd
	numerator := n1 * multiplier1 + n2 * multiplier2
	denomerator := d1 * multiplier1
	return numerator, denomerator
}