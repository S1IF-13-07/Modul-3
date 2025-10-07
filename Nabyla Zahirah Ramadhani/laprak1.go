package main

import "fmt"

func main (){
	var x, fx float64
	fmt.Print("Masukkan nilai fx: ")
	fmt.Scan(&fx)
	x = 2 / (fx-5) - 5
	fmt.Printf("Nilai x adalah: %.0f\n", x)
}