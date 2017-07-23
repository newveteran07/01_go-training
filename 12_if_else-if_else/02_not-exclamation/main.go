package main

import "fmt"

func main() {
     //! is shorthand for "not", so !true = "not true"
	if !true {
		fmt.Println("This did not run")
	}

	if !false {
		fmt.Println("This ran")
	}

}
