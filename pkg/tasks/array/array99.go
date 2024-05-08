package array

import "fmt"

func Array99() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	elementsFrequency := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
		if _, ok := elementsFrequency[ar[i]]; !ok {
			elementsFrequency[ar[i]] = 1
		} else {
			elementsFrequency[ar[i]]++
		}
	}

	tempAr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if elementsFrequency[ar[i]] <= 2 {
			tempAr = append(tempAr, ar[i])
		}
	}

	ar = tempAr
	m := len(ar)

	fmt.Println()
	fmt.Println("New size:", m)
	for i := 0; i < m; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}
