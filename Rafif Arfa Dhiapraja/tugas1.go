package main

import "fmt"

func main() {
	var fx float64

	fmt.Print("masukkan nilai fx = ")
	fmt.Scan(&fx)

	hasil := 2.0/(fx-5) - 5
	x := int(hasil)

	fmt.Printf("hasilnya adalah : %d", x)
}