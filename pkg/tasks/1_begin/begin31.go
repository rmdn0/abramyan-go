package begin

import "fmt"

func Begin31() {
	var tc, tf float32
	fmt.Print("T(f) = ")
	fmt.Scan(&tf)
	tc = (tf - 32) * 5 / 9
	fmt.Printf("\nT(c) = %.2f\n", tc)
}
