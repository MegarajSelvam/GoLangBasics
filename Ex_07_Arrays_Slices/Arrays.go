package Ex_07_Arrays_Slices

import "fmt"

// Array:
// Collection of items with same type
// Fixed Size
// Access via zero-based index
// len functino returns size of array
// copies refer to different underlying data

func Arrays() {
	// Array with size mentioned and items added
	grades := [3]int{97, 85, 83}
	fmt.Printf("Grades: %v\n", grades)

	// Array without mention size but items mentioned
	fruits := [...]string{"apple", "mango"}
	fmt.Printf("Fruits: %v\n", fruits)

	// Array with size but not defined array items
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Mega"
	students[1] = "Dileep"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #2: %v\n", students[1])
	// len is a function to identify the lenght of array
	fmt.Printf("Number of students: %v\n", len(students))
}

func ArraysOfArrays() {
	// Way 1
	var identityMatrix [3][3]int = [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1},
	}
	fmt.Println(identityMatrix)

	// Way 2
	var reverseIdentityMatrix [3][3]int
	reverseIdentityMatrix[0] = [3]int{0, 0, 1}
	reverseIdentityMatrix[1] = [3]int{0, 1, 0}
	reverseIdentityMatrix[2] = [3]int{1, 0, 0}
	fmt.Println(reverseIdentityMatrix)
}

func CopyingArray() {
	a := [...]int{1, 2, 3}
	b := a // It creates and new array and allocates new memory
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	c := &a // It creates the new array but it just points the same memory
	c[1] = 6
	fmt.Println(a)
	fmt.Println(c)
}
