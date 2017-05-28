package main

import "fmt"

const centemetersToInches float64 = 0.393701

func main() {
	var centemeters float64
	fmt.Print("Enter centemeters tall: ")
	fmt.Scan(&centemeters)
	inches := centemeters * centemetersToInches
	fmt.Println(centemeters, "centemeters is", inches, " inches.")
}
