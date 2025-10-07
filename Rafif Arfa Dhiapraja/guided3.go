package main

import "fmt"

func main() {
	var indo int 
	fmt.Print("Masukkan nilai uang dalam rupiah: ")
	fmt.Scan(&indo)
	var dolar = indo / 15000
	fmt.Print("Nilai uang dalam dolar: ", dolar)
}