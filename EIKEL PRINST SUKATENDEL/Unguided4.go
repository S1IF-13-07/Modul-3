package main

import "fmt"

func main() {
	var c float64
	fmt.Print("Suhu (°C): ")
	fmt.Scan(&c)

	fmt.Printf("Reamur: %.2f\nFahrenheit: %.2f\nKelvin: %.2f\n",
		c*4/5, c*9/5+32, c+273.15)
}