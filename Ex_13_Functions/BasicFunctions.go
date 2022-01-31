package Ex_13_Functions

import "fmt"

func BasicFunctions() {
	functionWithParameters("I am parameter function")

	for i := 0; i < 5; i++ {
		functionWithMultipleParameters("Message", i)
	}

	funcWithMultipleParamtersButCommonDataType("Hello", "Megay")

	passNameAsValueAndRef()

	sum(1, 2, 3, 4, 5)

	multiply("Multiply is: ", 1, 2, 3, 4, 5)

	greet := returnSomeValue("mega")
	fmt.Println(greet)

	multiply := returnLocalStackPointers(1, 2, 3)
	fmt.Println("Multiply is ", *multiply)

	fmt.Println(namedReturnFunc("megay"))

	d, err := returnMultipleValues(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println(d)
}

func passNameAsValueAndRef() {
	name := "mega"

	passNameAsValue(name)
	fmt.Println(name)

	passNameAsReference(&name)
	fmt.Println(name)
}

func passNameAsValue(name string) {
	fmt.Println(name)
	name = "subba"
	fmt.Println(name)
}

func passNameAsReference(name *string) {
	fmt.Println(*name)
	*name = "subba"
	fmt.Println(name)
}

func functionWithParameters(message string) {
	fmt.Println(message)
}

func functionWithMultipleParameters(message string, index int) {
	fmt.Println(message)
	fmt.Println(index)
}

func funcWithMultipleParamtersButCommonDataType(greeting, name string) {
	fmt.Println(greeting, name)
}

// Variadic Parameter
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
}

// Variadic parameter should always last parameter
// method should have only one Variadic parameter
// Received inside parameter as a slice
func multiply(message string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result *= v
	}
	fmt.Println(message, result)
}

func returnSomeValue(name string) string {
	return "Hello " + name
}

func returnLocalStackPointers(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result *= v
	}
	return &result
}

// Not used much in Go, as it leads to confusion
// Named return values
func namedReturnFunc(input string) (output string) {
	output = "Good Night" + input
	return
}

func returnMultipleValues(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
