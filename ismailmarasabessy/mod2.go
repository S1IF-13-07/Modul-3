package main

import (
	"fmt"
)

func main() {
	var indo int
	fmt.Print("Masukan nilai uang dalam rupiah: ")
	fmt.Scan(&indo)
	var dolar = indo / 15000
	fmt.Printf("nilai uang dalam dolar: %d ", dolar)
}
