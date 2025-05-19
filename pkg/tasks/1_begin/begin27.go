package begin

import "fmt"

func Begin27() {
	var a, temp float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	temp = a * a
	fmt.Printf("\nA^2 = %.2f\n", temp)
	temp = temp * temp
	fmt.Printf("A^4 = %.2f\n", temp)
	temp = temp * temp
	fmt.Printf("A^8 = %.2f\n", temp)
}
