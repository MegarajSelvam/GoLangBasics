package main

import (
	"golangbasics/Ex_01_Hello_World"
	"golangbasics/Ex_02_Variables"
	"golangbasics/Ex_03_Conversions"
	"golangbasics/Ex_04_Primitives"
	"golangbasics/Ex_05_Operators"
	"golangbasics/Ex_06_Constants"
)

func main() {
	Ex_01_Hello_World.PrintHelloWorld()

	Ex_01_Hello_World.WelcomeUser("megay")

	Ex_01_Hello_World.GuestureMessage()

	Ex_02_Variables.VariableDeclarationStyle()

	Ex_02_Variables.Shadowing()

	Ex_03_Conversions.NumericalConversions()

	Ex_03_Conversions.StringConversions()

	Ex_03_Conversions.StringToBytes()

	Ex_04_Primitives.BooleanDataTypes()

	Ex_04_Primitives.ComplexDataType()

	Ex_04_Primitives.NumberDataType()

	Ex_04_Primitives.TextDataType()

	Ex_05_Operators.ArithmeticOperators()

	Ex_05_Operators.BitShifting()

	Ex_05_Operators.BitWiseOperators()

	Ex_06_Constants.Constants()

	Ex_06_Constants.MixConstantAndVarialbe()

	Ex_06_Constants.ShadowingInConstants()

	Ex_06_Constants.EnumeratedConstants()

	Ex_06_Constants.EnumeratedExpressions()

}
