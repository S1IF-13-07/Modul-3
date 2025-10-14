package main

import "fmt"

func main() {
	var celsius float64

	// Input suhu dalam Celsius
	fmt.Print("Temperatur Celsius: ")
	fmt.Scanln(&celsius)

	// Konversi
	reamur := celsius * 4 / 5
	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273

	// Output
	fmt.Println("Derajat Reamur:", int(reamur))
	fmt.Println("Derajat Fahrenheit:", int(fahrenheit))
	fmt.Println("Derajat Kelvin:", int(kelvin))
}
