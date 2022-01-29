package Ex_03_Conversions

import (
	"fmt"
	"strconv"
)

func NumericalConversions() {
	var num1 int = 12
	var num2 float32 = float32(num1)
	fmt.Printf("%v %T \n", num2, num2)

	var num3 float32 = 12.34
	var num4 int = int(num3)
	// Converting lower data type to higher data type leads to loss of value
	fmt.Printf("%v %T \n", num4, num4)
}

func StringConversions() {
	num1 := 23

	// Problem of using this approach
	// conversion from int to string yields a string of one rune, not a string of digits
	/*
		var convertedStringWithoutstrconvPackage string = string(num1)
		fmt.Printf("%v %T \n", convertedStringWithoutstrconvPackage, convertedStringWithoutstrconvPackage)
	*/

	// We need to use this approach
	// Whenever we need to convert from other primitve data type to string, then we need to use strconv libray (https://pkg.go.dev/strconv)
	var convertedStringWithstrconvPackage string = strconv.Itoa(num1)
	fmt.Printf("%v %T \n", convertedStringWithstrconvPackage, convertedStringWithstrconvPackage)
}

func StringToBytes() {
	s := "this is a string"
	b := []byte(s)
	fmt.Printf("%v %T\n", b, b)
}
