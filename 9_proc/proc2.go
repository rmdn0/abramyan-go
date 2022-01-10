package main

import "fmt"

func main() {
	var a, b, c, d float64
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		powerA234(a, &b, &c, &d)
		fmt.Printf("%.2f %.2f %.2f\n", b, c, d)
	}
}

func powerA234(a float64, b *float64, c *float64, d *float64) {
	*b = a * a
	*c = *b * a
	*d = *c * a
}