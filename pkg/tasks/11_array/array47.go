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

	amount := 0
	var check bool
	for i := 0; i < n; i++ {
		check = true
		for j := 0; j < i; j++ {
			if ar[i] == ar[j] {
				check = false
				break
			}
		}
		if check {
			amount++
		}
	}

	fmt.Println(amount)
}
