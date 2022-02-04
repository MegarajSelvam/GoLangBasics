package Ex_10_Looping

import (
	"fmt"
)

func Looping() {
	simpleLooping()
	matrixLooping()
	scopeOfVariables()
	infiniteLoop()
	continueStatementInLoop()
	nestedLoop()
	breakWithLabel()
	loopOverSlice()
	loopOverMaps()
	loopOverString()
}

func simpleLooping() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func matrixLooping() {
	// In this case, we cannot use i++ or j++
	for i, j := 0, 1; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
}

func scopeOfVariables() {
	// i scoped function level
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(i) // we can use the i out side the loop

	// K is scoped only for "For loop"
	for k := 0; k < 5; k++ {
		fmt.Println(k)
	}
	// fmt.Println(k) // We cannot use k outside the loop

	l := 5
	for l < 10 {
		fmt.Println(l)
		l++
	}
}

func infiniteLoop() {
	i := 0

	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
}

func continueStatementInLoop() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func nestedLoop() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
		}
	}
}

func breakWithLabel() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break // This just break Inner loop.
			}
		}
	}

	// In order to break the entire looping, we can use label concept
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
}

func loopOverSlice() {
	s := []int{1, 2, 3}

	//  Key will give index and value will give actual value
	for key, value := range s {
		fmt.Println(key, value)
	}
}

func loopOverMaps() {
	statePopulations := map[string]int{
		"California": 112343,
		"Texas":      21356,
		"Florida":    323223,
		"New York":   412323,
	}

	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	// To get only key
	for k := range statePopulations {
		fmt.Println(k)
	}

	// To get only value
	for _, v := range statePopulations {
		fmt.Println(v)
	}
}

func loopOverString() {
	s := "Hello World"
	for k, v := range s {
		fmt.Println(k, v)
		fmt.Println(k, string(v))
	}
}
