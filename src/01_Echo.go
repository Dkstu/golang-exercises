package exercises

import (
	"fmt"
)

func init() {
	fmt.Println("Exercises_01 => Echo")
}

/**
 * Print & return SayHi
 **/
func Echo(text string) string {
	fmt.Println(text)
	return text
}
