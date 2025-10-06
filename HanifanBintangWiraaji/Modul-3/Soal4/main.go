package main

import "fmt"

func main() {
	var c int
	fmt.Print("Temperatur Celcius : ")
	fmt.Scanln(&c)
	var f int = c * 9 / 5 + 32
	var r int = ( c / 5 ) * 4
	var k int = c + 273
	fmt.Printf("Derajat Reamur : %v\n", r)
	fmt.Printf("Derajat Fahrenheit : %v\n", f)
	fmt.Printf("Derajat Kelvin : %v\n", k)
}