package main

import "fmt"

func main() {
	var bigNum int
	var smNum int

	fmt.Print("Please enter a large number: ")
	fmt.Scan(&bigNum)
	fmt.Print("Please enter a smaller number: ")
	fmt.Scan(&smNum)

	fmt.Println(bigNum, "%", smNum, "=", bigNum%smNum)
}
