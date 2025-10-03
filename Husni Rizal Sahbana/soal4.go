package main

import "fmt"

func main() {
	var c float64
	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&c)

	reamur := c * 4.0 / 5.0
	fahrenheit := c*9.0/5.0 + 32.0
	kelvin := c + 273.15

	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}
