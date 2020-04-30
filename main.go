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
	result, _ := exercises.Switch("Apple")
	fmt.Println(result)
}
