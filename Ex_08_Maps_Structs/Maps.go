package Ex_08_Maps_Structs

import "fmt"

// Collect Same data
// Multiple Assignment refer to the same data
func Maps() {
	// Way 1 - Literals
	statePopulations := map[string]int{
		"California": 112343,
		"Texas":      21356,
		"Florida":    323223,
		"New York":   412323,
	}
	fmt.Println(statePopulations)
	sp := statePopulations // This will affect the state of both the variable
	delete(sp, "Florida")
	fmt.Println(sp)
	fmt.Println(statePopulations)

	// Way 2 - Make Functinos
	districtPopulation := make(map[string]int)
	// districtPopulation := make(map[string]int, 10)
	districtPopulation = map[string]int{
		"DPI": 123445,
		"SM":  1243343,
		"CBT": 4562344,
	}
	fmt.Println(districtPopulation)
	fmt.Println(districtPopulation["DPI"])
	districtPopulation["DPI"] = 567890
	fmt.Println(districtPopulation["DPI"])
	fmt.Println(districtPopulation) // Order of map will get changed when we update any value

	delete(districtPopulation, "SM")
	fmt.Println(districtPopulation)
	fmt.Println(districtPopulation["SM"]) // It gives zero. Though we deleted the record still we are getting population as zero.

	// To Check the presence of item
	pop, ok := districtPopulation["SM"]
	fmt.Println(pop, ok) // 0 false
	// In above line, false indicates, this entry is not there in key value collection

	// To find the number of element
	fmt.Println(len(statePopulations))

	m := map[[3]int]string{} // Array is a valid map and slice is not a valid map
	fmt.Println(m)
}
