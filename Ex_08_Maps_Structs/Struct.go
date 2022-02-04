package Ex_08_Maps_Structs

import (
	"fmt"
	"reflect"
)

// Collects different data types
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func Struct() {
	// Way 1 - Use this approach
	doctor1 := Doctor{
		number:    3,
		actorName: "Mega",
		companions: []string{
			"Dileep",
			"Subba",
		},
	}

	// Way 2 - Don't use this approach
	// In future, if some one updated the strucutre of doctor struct it will be a problem
	doctor2 := Doctor{
		3,
		"Mega",
		[]string{
			"Dileep",
			"Subba",
		},
	}

	fmt.Println(doctor2)
	// 	fmt.Println(doctor2.actorName)
	fmt.Println(doctor1.actorName)
	fmt.Println(doctor1.companions[0])

	// Anonymous struct
	// Use case to generate JSON response
	anonymousDoctor := struct{ name string }{name: "Nithai"}
	fmt.Println(anonymousDoctor)

	anotherDoctor := anonymousDoctor
	anonymousDoctor.name = "Gaurav"
	fmt.Println(anonymousDoctor) // Nithai
	fmt.Println(anotherDoctor)   // Gaurav

	doctorPointToMemory := &anonymousDoctor
	doctorPointToMemory.name = "Pandit"
	fmt.Println(doctorPointToMemory) // Pandit
	fmt.Println(anonymousDoctor)     // Pandit

}

// Go doesn't support inheritance
// We have composition as a replacement. HAS A relationship

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   // We just embedding here not creating Animal Type (i.e., Animal Animal)
	SpeedKPH float32
	CanFly   bool
}

func Composition() {
	b := Bird{}
	b.Name = "EMU"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)      // {{EMU, Australia}, 48, false}
	fmt.Println(b.Name) // EMU

	anotherBird := Bird{
		Animal:   Animal{Name: "EMU", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(anotherBird)
}

// Tags can be added to struct fields to describe the fields
type SeaAnimals struct {
	Name   string `required max:"100"` // Tags
	Origin string
}

func TagsInStruct() {
	t := reflect.TypeOf(SeaAnimals{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
