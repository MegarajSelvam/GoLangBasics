package Ex_02_Variables

import "fmt"

// Different way to declare the variables outside the scope
/*
// Way 1: Grouping Variables
var (
	name1 string = "megay"
	job1  string = "programmer"
	age1  uint16 = 26
)

// Way 2: Declare the variable, type and assign the value
var name2 string = "Sujth"

// Way 3: Declare the variable and type and assign the value later
var name3 string
*/

//#############################################
// Note:

// name4 := "somerandomeName"
// This way we cannot declare the variable at package level

// If the variable name starts with Caps letter, then it can be accessed by another package

// If the varialbe name starts with small letter, then it is scoped with in the package

//#############################################

func VariableDeclarationStyle() {
	// Differnet way to declare varialbes inside scope
	// Way 1: Declare the variable, type and assign the value in single line
	var num1 int = 12

	// Way 2: Declare the variable and assign the value and then let Go decide the data type
	num2 := 13

	// Way 3: Declare the variable in one line and assign the value inside somewhere in program
	/*
		var num3 int
		num3 = 14
	*/

	// %v gives the actual value
	// %T gives the data type
	fmt.Printf("%v %T \n", num1, num1)
	fmt.Printf("%v %T \n", num2, num2)
	// fmt.Printf("%v %T", num3, num3)

	//#############################################
	// Note:
	// Grouping Variables is not possible inside the scope
	/*
		var (
			name1 string = "megay"
			job1  string = "programmer"
			age1  uint16 = 26
		)
	*/
	//#############################################
}

func Shadowing() {
	// Shadowing: Variable with inner most scope actually takes precedence
	// Example: I declared the variable named name2 at package level as well as method level
	// In that case, variable declared at inner most scope takes predence

	name2 := "Sonali"
	fmt.Println("Shadowing concept: ", name2)
}
