package main

import "fmt"

func main() {
	var t, h, m, s int
	for i := 0; i < 5; i++ {
		fmt.Print("T = ")
		fmt.Scan(&t)
		timeToHMS(t, &h, &m, &s)
		fmt.Printf("H = %d M = %d S = %d\n\n", h, m, s)
	}
}

func timeToHMS(t int, h, m, s *int) {
	*h = t / 3600
	*m = t % 3600 / 60
	*s = t % 60
}