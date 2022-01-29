package Ex_05_Operators

import "fmt"

func ArithmeticOperators() {
	a := 12
	b := 3

	fmt.Printf("Add: %v \n", a+b)
	fmt.Printf("Sub: %v \n", a-b)
	fmt.Printf("Multiply: %v \n", a*b)
	fmt.Printf("Division: %v \n", a/b)
	fmt.Printf("Remainder: %v \n", a%b)
}

func BitWiseOperators() {
	a := 10 //1010
	b := 3  // 0011

	// AND
	fmt.Println(a & b) // 0010
	// OR
	fmt.Println(a | b) // 1011
	// XOR
	fmt.Println(a ^ b) // 1001
	// AND NOT
	fmt.Println(a &^ b) // 0100
}

func BitShifting() {
	a := 8              // 2^3
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0
}
