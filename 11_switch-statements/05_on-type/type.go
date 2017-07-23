package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("String")
	case contact:
		fmt.Println("Contact")
	default:
		fmt.Println("unknown")

	}
}

func main() {
	SwitchOnType(42)
	SwitchOnType("McLeod")
	var t = contact{"Good to see you,", "Tim"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}
