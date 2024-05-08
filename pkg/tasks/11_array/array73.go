package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	var k, l int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("L = ")
	fmt.Scan(&l)

	distance := l - k
	step := 1

	for i := k; i < k+distance/2; i++ {
		step++
		fmt.Println(i, l, step, l-step)
		ar[i], ar[l-step] = ar[l-step], ar[i]
	}

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}
