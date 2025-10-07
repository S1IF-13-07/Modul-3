package main

import "fmt"

func main() {
	var celcius float64
	
	fmt.Print ("Temperatur celcius: ")
	fmt.Scan(&celcius)

	var fahreinheit int = int((celcius * 9/5) + 32)
	var reamur int = int(celcius * 4/5)
	var kelvin int = int(celcius + 273.15)
	
	fmt.Println("Derajat fahreinheit =", fahreinheit)
	fmt.Println("Derajat reamur = ", reamur)
	fmt.Println("Derajat kelvin = ", kelvin)

	

}