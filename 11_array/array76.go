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

	for i := 0; i < n; i++ {
		if i == 0 {
			if ar[i] > ar[i+1] {
				ar[i] = 0
				i++
			}
		} else if i < n-1 {
			if ar[i] > ar[i-1] && ar[i] > ar[i+1] {
				ar[i] = 0
				i++
			}
		} else if i == n-1 {
			if ar[i] > ar[i-1] {
				ar[i] = 0
				i++
			}
		}
	}

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}
