package main

import "fmt"

func main() {
	for i := 48; i <= 129; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
}
