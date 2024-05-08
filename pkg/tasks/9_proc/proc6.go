package main

import "fmt"

func main() {
	var k, c, s int
	for i := 0; i < 5; i++ {
		fmt.Print("K = ")
		fmt.Scan(&k)
		digitCountSum(k, &c, &s)
		fmt.Printf("C = %d S = %d\n", c, s)
	}
}

func digitCountSum(k int, c, s *int) {
	*c = 0
	*s = 0
	for k > 0 {
		*s += k % 10
		k /= 10
		*c++
	}
}