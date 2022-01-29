package Ex_04_Primitives

import "fmt"

func BooleanDataTypes() {
	// Default value: false
	/*

		var mybool1 bool = false
		fmt.Printf("mybool1 is %v", mybool1)

		var mybool2 bool
		fmt.Printf("mybool2 is %v", mybool2)
		mybool2 = true
		fmt.Printf("mybool2 is %v", mybool2)

		mybool3 := false
	*/
	status1 := 1 == 1
	status2 := 1 != 2
	status3 := 1 == 12

	fmt.Printf("Status1 is %v \n", status1)
	fmt.Printf("Status2 is %v \n", status2)
	fmt.Printf("Status3 is %v \n", status3)
}

func NumberDataType() {

	//Signed Integer: int8, int16, int32 and int64
	//Unsigned Integer: uint8, uint16 and uint32
	// Available Arithmetic Operations: + - * / and %
	// Note:
	// We cannot mix types
	// Example: int16 + int32 give error
	// Default value: 0

	// Float: float32, float64
	// Default value: 0
	// Available Arithemtic Operations: + - * /
	// Float Literal styles
	// 1) Decimal 3.14
	// 2) Exponentital 13e18 or 2E10
	// 3) Mixed 13.7e12

	num1 := 23
	fmt.Printf("%v %T \n", num1, num1)

	num2 := -23
	fmt.Printf("%v %T \n", num2, num2)
}

func ComplexDataType() {
	// complex: complex64, complex128
	// Default value: 0 + 0i
	// Available Arithemtic Operations: + - * /
	// Built In Functions
	// complex: make complex number from 2 float
	// real:get real part as a float
	// imag: get imaginary part as a float
	var compNumber1 complex64 = 1 + 2i
	fmt.Printf("%v %T \n", compNumber1, compNumber1)

	var compNumber2 complex64 = complex(16, 23)
	fmt.Printf("%v %T \n", compNumber2, compNumber2)

	var compNumber3 complex64 = 2i
	fmt.Printf("%v %T \n", real(compNumber3), real(compNumber3))
	fmt.Printf("%v %T \n", imag(compNumber3), imag(compNumber3))
}

func TextDataType() {
	// string: represents UTF8 character
	// Immutable
	// Can be concatenated with + symbol
	// Can be converted to []byte
	s := "This is a string"
	fmt.Printf("%v %T \n", s, s)

	// rune: represents UTF32 character
	// rune is an alias for int32
	// Special method normally required to process
	// Eg: strings.Reader#ReadRune

	// rune will be declared in single quotes
	r := 'r'         // Way 1 to declare rune
	var q rune = 'q' // Way 2 to declare rune
	fmt.Printf("%v %T\n", r, r)
	fmt.Printf("%v %T\n", q, q)
}
