package main

import "fmt"

func main(){
	var celcius float64

	fmt.Printf("Temperatur celcius: ")
	fmt.Scan(&celcius)
	
	reamur := (4.0 / 5.0) * celcius
	fahrenheit := (9.0 / 5.0) * celcius + 32
	kelvin := celcius + 273.15

	fmt.Printf("Derajat Reamur: %.f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.f\n", kelvin)
}