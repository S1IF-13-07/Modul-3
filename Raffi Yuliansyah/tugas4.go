package main

import "fmt"

func main() {
	var celcius, reamur, fahrenheit, kelvin float64
	fmt.Print("Input, derjat temperatur Celcius = ")
	fmt.Scanln(&celcius)
	reamur = (celcius * 4) / 5
	fahrenheit = ((celcius * 9) / 5) + 32
	kelvin = celcius + 273
	fmt.Println("Output, Derajat Reamur : ", reamur)
	fmt.Println("Output, Derajat Fahrenheit : ", fahrenheit)
	fmt.Println("Output, Derajat Kelvin : ", kelvin)
}