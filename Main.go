package main

import (
	"golangbasics/Ex_01_Hello_World"
	"golangbasics/Ex_02_Variables"
	"golangbasics/Ex_03_Conversions"
	"golangbasics/Ex_04_Primitives"
	"golangbasics/Ex_05_Operators"
)

func main() {
	// Printing Hello World
	Ex_01_Hello_World.PrintHelloWorld()

	// Welcome User
	Ex_01_Hello_World.WelcomeUser("megay")

	// Guesture Message
	Ex_01_Hello_World.GuestureMessage()

	// How to declare variables and thier scopes
	Ex_02_Variables.VariableDeclarationStyle()

	// Shadowing Concept
	Ex_02_Variables.Shadowing()

	// Convert from one numeric type to another
	Ex_03_Conversions.NumericalConversions()

	// Convert primitive data type to string
	Ex_03_Conversions.StringConversions()

	Ex_03_Conversions.StringToBytes()

	Ex_04_Primitives.BooleanDataTypes()

	Ex_04_Primitives.ComplexDataType()

	Ex_04_Primitives.NumberDataType()

	Ex_04_Primitives.TextDataType()

	Ex_05_Operators.ArithmeticOperators()

	Ex_05_Operators.BitShifting()

	Ex_05_Operators.BitWiseOperators()

}
