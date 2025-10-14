package main

import "fmt"

func main() {
	var fx float64
	fmt.Print("Masukkan nilai fx: ")
	fmt.Scan(&fx)
	fmt.Printf("Nilai X: %.2f\n", 2/(fx-5)-5)
}