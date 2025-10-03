package main

import "fmt"

func main(){
	var fx float64
	fmt.Print("Masukkan Nilai fx: ")
	fmt.Scan(&fx)
	nilaiX := (2/(fx-5)) - 5
	fmt.Printf("Nilai X nya Adalah: %.0f", nilaiX)
}