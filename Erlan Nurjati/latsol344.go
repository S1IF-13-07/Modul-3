package main

import "fmt"

func main() {
	var celsius float64
	fmt.Print("Masukkan celcius: ")
	fmt.Scan(&celsius)

	reamur := celsius * 4.0 / 5.0
	fahrenheit := (celsius * 9.0 / 5.0) + 32
	kelvin := celsius + 273

	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}
