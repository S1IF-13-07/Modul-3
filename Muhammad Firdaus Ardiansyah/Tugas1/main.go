package main

import "fmt"

func main(){
	var fx float64
	fmt.Printf ("Masukkan nilai x: ")
	fmt.Scan(&fx)

	hasil := (2 / (fx - 5 )) - 5
	x := int(hasil)
	fmt.Printf ("nilai dari x adalah %d", x)
}
