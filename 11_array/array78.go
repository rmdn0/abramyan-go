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

	var avg, prev float64
	for i := 0; i < n; i++ {
		if i == 0 {
			avg = (ar[i] + ar[i+1]) / 2
			prev = ar[i]
			ar[i] = avg
		} else if i < n-1 {
			avg = (prev + ar[i] + ar[i+1]) / 3
			prev = ar[i]
			ar[i] = avg
		} else {
			avg = (prev + ar[i]) / 2
			ar[i] = avg
		}
	}

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%.2f ", ar[i])
	}
	fmt.Println()
}
