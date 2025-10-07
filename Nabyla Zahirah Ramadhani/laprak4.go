package main

import "fmt"

func main (){
	var celcius, fahrenheit float64
	fmt.Print("Masukan celcius: ")
	fmt.Scan(&celcius)
	fmt.Print("Masukan fahrenheit: ")
	fmt.Scan(&fahrenheit)
	reamur := celcius * 4 / 5
	kelvin := (fahrenheit + 459.67 ) * 5 / 9

	fmt.Println("Temperatur celcius: ", celcius)
	fmt.Println("Derajat fahrenheit: ", fahrenheit)
	fmt.Println("Derajat reamur: ", reamur)
	fmt.Printf("Derajat kelvin: %.0f\n", kelvin)
}