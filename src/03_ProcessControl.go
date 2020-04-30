package exercises

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("Exercises_03 => Process Control")
}

// If > 18 = 20, < 18 & < 15 = 10, other = 1
func If(age int) int {
	var result int
	if age > 18 {
		result = 20
	} else if age < 15 {
		result = 10
	} else {
		result = 10
	}
	return result
}

// For Sum 1 ~ 9
func For() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var sum = 0
	for _, val := range ints {
		sum += val
		fmt.Print(val)
	}
	fmt.Println("\nSum:", sum)
}

// Goto Print Safe!
func Goto() {
	goto Here
	fmt.Println("Boom!")
Here:
	fmt.Println("Safe!")
}

// Switch Case Apple Banana Orange
func Switch(name string) (string, error) {
	var result string
	var err error
	switch name {
	case "Apple":
		result = "Apple match"
	case "Banana":
		result = "Banana match"
	case "Orange":
		result = "Orange match"
	default:
		err = errors.New("noting match")
	}
	return result, err
}
