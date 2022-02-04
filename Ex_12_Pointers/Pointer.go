package Ex_12_Pointers

import "fmt"

/*
// Creating Pointer
- Pointer types use an asterisk (*) as a prefix to type pointed to
	- *int - a pointer to an integer
- Use the addressof operator (&) to get address of variable

// Dereferencing Pointers
- Dereference a pointer by preceding with an asterisk(*)
- Complex types (eg: structs) are automatically dereferenced

// Types with internal pointers
- All assignment operations in Go are copy operations
- Slices and maps contain internal pointers, so copies point to same underlying data
*/

func Pointers() {
	basicExample()
	pointerArithmetic() // Pointer arithemtic is not allowed in Go lang
	pointerInStruct()
	pointerInArray()
	pointerWithSlice()
	pointerWithMap()
}

func basicExample() {
	// Without Pointer
	a := 12
	b := a // this just created another variable in another memory location. It is not pointing the same memory loction of a. so changing anything to b will not affect a and vice versa
	fmt.Println(a, b)

	a = 27
	fmt.Println(a, b)

	// With Pointer
	var c int = 21
	var d *int = &a    // d is holding the memory location of c
	fmt.Println(c, d)  // value of c and memory loction of d
	fmt.Println(&c, d) // memory loction of c and memory location of d
	fmt.Println(c, *d) // value of c and value of d
	c = 25
	fmt.Println(c, *d)
	*d = 14
	fmt.Println(c, *d)
}

func pointerArithmetic() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[2]
	// tyring to do arithmetic operation
	// d := &a[2]-4 // It is not allowed in Go due to unsafe issue

	fmt.Printf("%v %p %p\n", a, b, c)
}

func pointerInStruct() {
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	var mds *myStruct
	fmt.Println(mds) // Inital value of pointer is Nil
	mds = new(myStruct)
	(*mds).foo = 42
	fmt.Println(mds)        // &{42}
	fmt.Println((*mds).foo) // 42
	fmt.Println(mds.foo)
}

type myStruct struct {
	foo int
}

func pointerInArray() {
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 42
	// Changing a will not affect b
	fmt.Println(a, b)
}

func pointerWithSlice() {
	a := []int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 42
	// Changing a will affect b
	fmt.Println(a, b)
}

func pointerWithMap() {
	a := map[string]string{"foo": "bar", "baz": "buz"}
	b := a
	fmt.Println(a, b)
	a["foo"] = "qux"
	// Changing a will affect b
	fmt.Println(a, b)
}
