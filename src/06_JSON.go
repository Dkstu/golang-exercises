package exercises

import (
	"encoding/json"
	"fmt"
)

func JSON() {
	person := Man{
		Name: "Helen",
		Age:  20,
	}

	fmt.Println(person)
	out, err := json.Marshal(person)
	if err != nil {
		fmt.Println("json marsha1 error")
	}
	fmt.Printf("%s\n", out)

	var m Man
	err = json.Unmarshal([]byte(out), &m)
	if err != nil {
		fmt.Println("json unmarshal error: ", err)
	}
	fmt.Println(m)
}
