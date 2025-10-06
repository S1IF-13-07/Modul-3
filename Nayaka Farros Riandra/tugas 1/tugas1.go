package main

import "fmt"

func main() {
	var fx float64

	fmt.Print("masukkan nilai x = ")
	fmt.Scan(&fx)

	hasil := 2/(fx-5) - 5
	x := int(hasil)

	fmt.Printf("hasil fx adalah = %d", x)
}