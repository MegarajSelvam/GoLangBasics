package Ex_09_If_And_Switch_Statements

import "fmt"

func SwitchExample() {
	basicSwitchStatement()
	multipleStatementInCase(5)
	simpleOperationInSwitch()
	tagLessSwitchStatement()
	switchWithInterface()
}

func basicSwitchStatement() {
	switch 2 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Neither One nor Two")
	}
}

func multipleStatementInCase(num int) {
	switch num {
	case 2, 4, 6, 8, 10:
		fmt.Println("Number is even")
	case 1, 3, 5, 7, 9:
		fmt.Println("Number is odd")
	default:
		fmt.Println("Number is bigger than 10 or less than or equal to 0")
	}
}

func simpleOperationInSwitch() {
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("One, Five or Ten")
	case 2, 4, 6:
		fmt.Println("Two, Four or Six")
	default:
		fmt.Println("Another Number")
	}
}

func tagLessSwitchStatement() {
	i := 10
	switch {
	case i <= 10:
		fmt.Println("Less than or equal to ten")
	case i <= 20:
		fmt.Println("Less than or equal to Twenty")
	default:
		fmt.Println("Greater than Twenty")
	}
	/*
		// Output:
		Less than or equal to ten
	*/

	switch {
	case i <= 10:
		fmt.Println("Less than or equal to ten")
		fallthrough // Continue execute other statements
	case i <= 20:
		fmt.Println("Less than or equal to Twenty")
	default:
		fmt.Println("Greater than Twenty")
	}
	/*
		// Output:
		Less than or equal to ten
		Less than or equal to Twenty
	*/
}

func switchWithInterface() {
	var i interface{} = 1
	// var i interface{} = [3]int{}

	switch i.(type) {
	case int:
		fmt.Println("i is integer type")
		break // Breaking out Early
		fmt.Println("Due to break it will not get execute")
	case float64:
		fmt.Println("i is float type")
		fmt.Println("It will get execute. If you dont want to execute this line add break before this line like previous case")
	case string:
		fmt.Println("i is string type")
	default:
		fmt.Println("i is another type")
	}
}
