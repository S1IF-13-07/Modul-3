package main

import "fmt"

func main() {
	var celcius float64
	fmt.Print("Masukkan suhu dalam derajat celcius = ")
	fmt.Scanln(&celcius)

	fahrenheit := (celcius * 9.0 / 5.0) + 32.0
	Reamur := celcius * 4 / 5
	kelvin := (fahrenheit + 459.67) * 5 / 9

	fmt.Printf("Derajat Reamur = %.2f°Re\n", Reamur)
	fmt.Printf("Derajat Fahrenheit = %.2f°F\n", fahrenheit)
	fmt.Printf("Derajat Kelvin = %.2fK\n", kelvin)
}
