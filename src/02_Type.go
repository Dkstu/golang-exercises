package exercises

import "fmt"

func init() {
	fmt.Println("Exercises_02 => Type")
}

// Type
func Type() {
	// boolean
	var status bool = true
	if status {
	}

	// string
	// var str string = "Hello World"
	var str = "Hello World"
	if str == "Hello World" {
	}

	// unit8 => unsign integer 8-bit, range 0~255
	var uint8 uint8 = 255
	if uint8 == 255 {
	}

	// float32
	var float32 float32 = 0.1234567
	if float32 == 0.1234567 {
	}

	// var byte []byte = []byte{3, 3}
	var byte = []byte("Life is Good")
	fmt.Println(string(byte), ";", string(byte[8:12]), string(byte[5:7]), string(byte[:4]))

	// Array static length
	var arrStrings [2]string
	arrStrings[0] = "Hello"
	arrStrings[1] = "World!"
	fmt.Println(arrStrings)

	// Slice dynamic length
	var sliStrings = []string{"Wash", "Wash"}
	sliStrings = append(sliStrings, "Wash", "Hand")
	fmt.Println(sliStrings)

	// Map
	var m map[int]string
	m = make(map[int]string)
	m[1] = "Amy"
	m[2] = "John"
	_, exist := m[3]
	if !exist {
		fmt.Println("m[3] not exist")
	}

	fmt.Println("Before:", m)
	delete(m, 1)
	fmt.Println("After:", m)

	// Struct
	fmt.Println(Person{"Amy", 123})
	person := new(Person)
	person.name = "Glenn"
	person.age = 999
	fmt.Println(person)

}
