package main

import (
	"fmt"
)

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

//the fist thing it prints is a "slice" its a collection of numbers
//The second thing printed is tye type
// the third thing printed is the average
//You can see just the average by removing the other two "Prints"