package main

import "fmt"

func main() {
	var celsius float64

	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9.0 / 5.0) + 32

	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
}
