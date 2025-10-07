package main

import "fmt"

func main() {
	var c float64
	fmt.Print("Masukkan temperatur dalam Celsius: ")
	fmt.Scan(&c)

	f := (9.0/5.0)*c + 32
	r := (4.0 / 5.0) * c
	k := c + 273

	fmt.Println("Derajat Reamur:", r)
	fmt.Println("Derajat Fahrenheit:", f)
	fmt.Println("Derajat Kelvin:", k)
}