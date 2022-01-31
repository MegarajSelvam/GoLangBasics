package Ex_13_Functions

import "fmt"

func AdvancedFunctions() {

	anonymousMethodExample()

	// Function as a Type
	functionAsAVariablesExample()

	// Methods: Function that executes in context of a type
	// Function is executing in unknown context
	functionIntoAMethodExample()
}

func anonymousMethodExample() {
	// Anonymous functions
	// It will immediate invoke after execution
	// It can be assigned to a variable or passed as an argument to a function
	func() {
		msg := "Hello Anonymous function"
		fmt.Println(msg)
	}()

	// some use case for Anonymous function
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func functionAsAVariablesExample() {
	// Functions as a variables
	// Way 1
	var fun func() = func() {
		fmt.Println("Function as a variables")
	}
	fun()
	// Way 2
	f := func() {
		fmt.Println("Function as a variables")
	}
	f()

	// Some use case for Function as a variables
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	}
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
}

type greeter struct {
	greeting string
	name     string
}

func functionIntoAMethodExample() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("The new name is: ", g.name)

	g.greetPointer()
	fmt.Println("The new name is: ", g.name)
}

//
func (g greeter) greet() {
	// here we are working on the copy of the greeter object. not the greeter object itself
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

func (g *greeter) greetPointer() {
	// here we are working on the copy of the greeter object. not the greeter object itself
	fmt.Println(g.greeting, g.name)
	g.name = ""
}
