package Ex_09_If_And_Switch_Statements

import "fmt"

func IfStatementsInGo() {
	simpleIfStatement()
	ifStatementForMaps()
	comparisionOperators()
	logicalOperators()
	ifAndElseStatement()
}

func simpleIfStatement() {
	if true {
		fmt.Println("The test is true")
	}

	/* Not a valid syntax
	// Even it is a single line block, we have to use curly braces
	if true
		fmt.Println("The test is true")
	*/
}

func ifStatementForMaps() {
	statePopulations := map[string]int{
		"California": 112343,
		"Texas":      21356,
		"Florida":    323223,
		"New York":   412323,
	}

	// Here pop variable hold populatino for florida
	// ok variable hold florida state is avaiable in collectin or not
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	// We cannot access the pop variable here. Because it is defined inside if block scope
	// fmt.Println(pop);
}

func comparisionOperators() {
	number := 50
	guess := 30

	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("Too High")
	}
	if guess == number {
		fmt.Println("You got it")
	}

	fmt.Println(number <= guess, number >= guess, number != guess)
}

func logicalOperators() {
	// Short Circuit: Use &&, ||
	// Long Circuit: Use &, |
	number := 50
	guess := 30

	// OR ||
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	}

	// And &&
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too High")
		}
		if guess == number {
			fmt.Println("You got it")
		}
	}

	fmt.Println(true)
	// Not !
	fmt.Println(!true) // false
}

func ifAndElseStatement() {
	num := 5
	if num > 5 {
		fmt.Println("The number is greater than 5")
	} else if num < 5 {
		fmt.Println("The number is lesser than 5")
	} else {
		fmt.Println("The number is equal to 5")
	}
}
