package main

import "fmt"

func main() {
	var x, y float64
	for i := 0; i < 3; i++ {
		fmt.Printf("x%d = ", i + 1)
		fmt.Scan(&x)
		fmt.Printf("y%d = ", i + 1)
		fmt.Scan(&y)
		fmt.Printf("\nQuarter(%.2f, %.2f) = %d\n", x, y, quarter(x, y))
	}
}

func quarter(x, y float64) int {
	quarter := 0
	switch {
	case x > 0:
		switch {
		case y > 0:
			quarter = 1
		case y < 0:
			quarter = 4
		}
	case x < 0:
		switch {
		case y > 0:
			quarter = 2
		case y < 0:
			quarter = 3
		}
	}
	return quarter
}