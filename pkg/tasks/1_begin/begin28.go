package begin

import "fmt"

func Begin28() {
	var a float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	temp := a * a
	fmt.Printf("\nA^2 = %.2f\n", temp)
	temp2 := temp * a
	fmt.Printf("A^3 = %.2f\n", temp2)
	temp *= temp2
	fmt.Printf("A^5 = %.2f\n", temp)
	temp2 = temp * temp
	fmt.Printf("A^10 = %.2f\n", temp2)
	temp *= temp2
	fmt.Printf("A^15 = %.2f\n", temp)
}
