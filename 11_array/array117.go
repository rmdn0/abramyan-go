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

	tempAr := []int{}
	var count int
	for i := 0; i < n; i++ {
		tempAr = append(tempAr, 0)
		count = 0
		for i+count < n-1 && ar[i+count] == ar[i+count+1] {
			count++
		}
		tempAr = append(tempAr, ar[i:i+count+1]...)
		i += count
	}
	ar = tempAr
	n = len(ar)
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}
