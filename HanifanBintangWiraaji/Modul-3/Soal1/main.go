package main

import "fmt"

func main() {
	var fx float64
	fmt.Print("Masukan x = ")
	fmt.Scanln(&fx)
	var x float64 = ( 2 / ( fx - 5 )) - 5
	var hasil int = int(x)
	fmt.Print(hasil)
}