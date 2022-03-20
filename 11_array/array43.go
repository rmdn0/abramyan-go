package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	distinctElementsAmount := 1

	for i := 0; i < n-1; {
		if ar[i] == ar[i+1] {
			for i < n-1 && ar[i] == ar[i+1] {
				i++
			}
		} else {
			distinctElementsAmount++
			i++
		}
	}

	fmt.Println(distinctElementsAmount)
}
