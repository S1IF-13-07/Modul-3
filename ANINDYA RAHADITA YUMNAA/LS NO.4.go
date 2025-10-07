package main

import "fmt"

func main () {
	var celcius float64

	fmt.Print("Temperatur Celcius: ")
	fmt.Scanln(&celcius)

	reamur := celcius * 4/5
	fahrenheit := (celcius * 9/5) + 32
	kelvin := celcius + 273

	fmt.Println("Derajat Reamur: ", reamur)
	fmt.Println("Derajat Fahrenheit: ", fahrenheit)
	fmt.Println("Derajat Kelvin", kelvin)
}