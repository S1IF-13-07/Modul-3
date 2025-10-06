package main

import "fmt"

func main() {
	var c float64

	fmt.Printf("masukkan temperatur celcius = ")
	fmt.Scan(&c)

	reamur := (4/5)*c
	fahrenheit := (9/5) * c + 32
	kelvin := c + 273.15

	fmt.Println("derajat reamur = ", reamur)
	fmt.Println("derajat fahrenheit = ", fahrenheit)
	fmt.Println("derajat kelvin = ", kelvin)
}