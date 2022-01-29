package Ex_01_Hello_World

import "fmt"

func PrintHelloWorld() {
	// Print - Printing Hello World without new line
	fmt.Print("Hello World")
	fmt.Print("\n") // Printing new line explicitly
}

func WelcomeUser(userName string) {
	// Printf - Formatted Print. It will not add any new line at the end
	// so we need to add line explicitly at the endby adding "\n"
	fmt.Printf("Welcome %v to learn Go Lang\n", userName)
}

func GuestureMessage() {
	// Declaring a variable named "guestureMessage"
	var guestureMessage string = "Have a nice day!!!"
	// Println - Print with new line at the end
	fmt.Println("Hello ", guestureMessage)
}
