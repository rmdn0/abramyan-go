package array

import "fmt"

func Array63() {
	const n = 5
	a := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	b := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}

	c := make([]float64, n*2, n*2)
	var i1, i2 int
	for i := 0; i < n*2; {
		if i1 == n {
			for ; i2 < n; i2++ {
				c[i] = b[i2]
				i++
			}
		}
		if i2 == n {
			for ; i1 < n; i1++ {
				c[i] = a[i1]
				i++
			}
		}
		if i == n*2 {
			break
		}
		if a[i1] < b[i2] {
			for i1 < n && a[i1] < b[i2] {
				c[i] = a[i1]
				i++
				i1++
			}
		} else if a[i1] > b[i2] {
			for i2 < n && a[i1] > b[i2] {
				c[i] = b[i2]
				i++
				i2++
			}
		} else {
			i++
		}
	}

	fmt.Println()
	for i := 0; i < n*2; i++ {
		fmt.Print(c[i], " ")
	}
	fmt.Println()
}
