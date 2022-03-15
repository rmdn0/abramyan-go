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

	fmt.Println(isPermutation(ar))
}

func isPermutation(ar []int) int {
	target := len(ar)
	n := target
	for i := 0; i < n; i++ {
		if ar[i] <= n {
			target--
		} else {
			return i + 1
		}
		if i == n-1 {
			break
		}
		for j := i + 1; j < n; j++ {
			if ar[i] == ar[j] {
				return j + 1
			}
		}
	}
	return target
}
