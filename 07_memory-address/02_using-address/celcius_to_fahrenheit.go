package main

import "fmt"

const celciustofahrenheit float64 = 1.8

func main() {
	var celcius float64
	fmt.Print("celcius in temp: ")
	fmt.Scan(&celcius)
	fahrenheit := celcius*celciustofahrenheit + 32
	fmt.Println(celcius, "celcius is", fahrenheit, "fahrenheit.")
}
