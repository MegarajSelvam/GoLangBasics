package Ex_06_Constants

import (
	"fmt"
	//"math"
)

const shadow int = 14

func Constants() {
	// Typed Constants
	const myConst1 int = 32

	// const myConst2 float64 = math.Sin(1.57)
	// const value need to assign when it is declared. not at compile time
	// so it gives error
	// Value must be calculable at compile time

	fmt.Printf("%v %T \n", myConst1, myConst1)
}

func ShadowingInConstants() {
	// Ideally we should not apply shadowing for constant
	// It violates the constant principle
	const shadow int = 12
	fmt.Println("Sahdowing in constants: ", shadow)
}

func MixConstantAndVarialbe() {
	// Type 1: Typed Constants
	const num1 int = 12
	var num2 int32 = 14

	// Compiler: int + int32
	// we cannot add num1 + num2 directly as it differs data types. so need explicit conversion to achieve num1 + num2
	// Can interpolate with same type
	fmt.Println(num1 + int(num2))

	// Type 2: Untyped Constants
	const num3 = 12
	var num4 int64 = 14
	// Compiler: 12 + int64
	// when the constant is declared without data type, then we dont required any conversion.
	// Can interpolate with similar type
	fmt.Println(num3 + num4)

}
