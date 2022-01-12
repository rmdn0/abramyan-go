package main

import "fmt"

func main() {
	var n int
	for i := 0; i < 5; i++ {
		fmt.Print("N = ")
		fmt.Scan(&n)
		fmt.Printf("%d\n\n", fib(n))
	}
}

func fib(n int) int {
	f1 := 1
	f2 := 1
	temp := 0
	for i := 3; i <= n; i++ {
		temp = f1
		f1 = f2
		f2 += temp
	}
	return f2
}