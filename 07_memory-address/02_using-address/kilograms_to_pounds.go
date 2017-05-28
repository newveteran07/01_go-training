package main

import "fmt"

const kilogramsToPounds float64 = 2.20462

func main() {
	var kilograms float64
	fmt.Print("Kilograms in weight: ")
	fmt.Scan(&kilograms)
	pounds := kilograms * kilogramsToPounds
	fmt.Println(kilograms, "kilograms is", pounds, "pounds.")
}
