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
