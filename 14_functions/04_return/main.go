package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

//greet is the name of the function
//(fname, lname string) are the paramiters
//string is the return
// so func *name* (paramiters) return