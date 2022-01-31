package Ex_07_Arrays_Slices

import "fmt"

// Slices:
// Backed by array
// make function (first parameter: data type; second parameter: length; third Parameter: capacity)
// len function returns the length of slice
// cap function return length of underlying arrays
// append function to add elements to slice. May cause expensive copy operation if underlying array is too small
// Copies refer to same underlying array

func Slices() {
	a := []int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}

func DifferentWaysOfCreatingSlices() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // slice of all elements
	c := a[3:]  // slice from the 4th element to end
	d := a[:6]  // slice first 6 element
	e := a[3:6] // slice the 4th, 5th and 6th elements

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func CreateSliceWithMakeBuiltInFunc() {
	a := make([]int, 3)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	b := make([]int, 3, 100)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(b))
	fmt.Printf("Capacity: %v\n", cap(b))
}

func AddingElementToSlice() {
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))   // value is 1
	fmt.Printf("Capacity: %v\n", cap(a)) // capacity is 2

	a = append(a, 2, 3, 4, 5)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))   // value is 5
	fmt.Printf("Capacity: %v\n", cap(a)) // capacity is 8

	a = append(a, []int{2, 3, 4, 5}...)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}

func StackOperationInSlices() {
	// Stack Operations: Push and Pop operations

	// To Remove the element at beginning
	a := []int{1, 2, 3, 4, 5}
	b := a[1:]
	fmt.Println(b)

	// To Remove the element at End
	a = []int{1, 2, 3, 4, 5} // assigning back all values for demo
	c := a[:len(a)-1]
	fmt.Println(c)

	// To Remove the element at Middle
	a = []int{1, 2, 3, 4, 5} // assigning back all values for demo
	fmt.Println(a)
	d := append(a[:2], a[3:]...)
	fmt.Println(d)
	fmt.Println(a)
}
