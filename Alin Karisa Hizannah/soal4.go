package main

import "fmt"

func main() {
	var celsius, reamur, fahrenheit, kelvin float64

	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celsius)

	// Konversi suhu
	reamur = celsius * 4.0 / 5.0
	fahrenheit = (celsius * 9.0 / 5.0) + 32.0
	kelvin = celsius + 273.0

	// Output hasil
	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}
