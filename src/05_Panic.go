package exercises

import (
	"errors"
	"fmt"
)

// CallPanic .
func CallPanic() {
	var err = errors.New("Panic QQ")

	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3")
	panic(err)
}

// CallPanicWithRecover .
func CallPanicWithRecover() {
	var err = errors.New("Panic QQ")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("I am recover.", err)
		}
	}()
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3")
	panic(err)
}
