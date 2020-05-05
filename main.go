package main

import (
	"fmt"
	exercises "golang-exercises/src"
)

func main() {
	// Echo
	exercises.Echo("Hello World")

	// Type
	exercises.Type()

	// ProcessControl
	exercises.For()
	exercises.Goto()
	result, error := exercises.Switch("Apple")
	if error != nil {
		fmt.Println(result)
	}
	fmt.Println(result)

	// Slice
	result1, result2 := exercises.CopySlice([]string{"A", "B"})
	fmt.Println(result1, result2)

	// Panic
	// exercises.CallPanic()
	exercises.CallPanicWithRecover()

}
