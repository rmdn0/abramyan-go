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

	var k int
	fmt.Print("K = ")
	fmt.Scan(&k)

	sumKthValueToEachElement(ar, k)

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%.2f ", ar[i])
	}
	fmt.Println()
}

func sumKthValueToEachElement(ar []float64, k int) {
	diff := ar[k-1]
	for i := 0; i < len(ar); i++ {
		ar[i] += diff
	}
}
