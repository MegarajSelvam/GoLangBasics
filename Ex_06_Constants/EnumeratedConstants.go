package Ex_06_Constants

import "fmt"

const a = iota

// Type 3: Enumerated Constant
const (
	// iota start will be zero, then it will add + 1 for every next iota defined in the block. Value of the iota is scopoed to the block
	b = iota
	c = iota
	d // even though we didn't mention it as iota, go will consider it as iota because we declared b and c as iota. so, go will consider d also as an iota
)

const (
	// errorSpecialist = iota // value = 0
	_               = iota // Ignore first value by assigning to blank identifier
	catSpecialist          // value = 1
	dotSpecialist          // value = 2
	snakeSpecialist        // value = 3
)

const (
	Mon = iota // value = 0
	Tue        // value = 1
	Wed        // value = 2
	Thu        // value = 3
	Fri        // value = 4
	Sat        // value = 5
	Sun        // value = 6
)

const (
	_       = iota + 5 // value
	target1            // value = 6
	target2            // value = 7
)

const (
	e = iota
)

func EnumeratedConstants() {
	fmt.Printf("%v %T \n", a, a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	var speciallistType int
	fmt.Printf("%v \n", speciallistType == catSpecialist)
	fmt.Println(target2)
}

// Type 4: Enumerated Expression
const (
	isAdmin = 1 << iota
	isHeadQuarters
	canSeeFinancials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
)

func EnumeratedExpressions() {
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v", isAdmin&roles == isAdmin)
}
