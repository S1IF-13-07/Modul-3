package main

import (
	"fmt"
)

func main() {

	var indo int
	fmt.Print("Masukan uang rupiah antum bray: ")
	fmt.Scan(&indo)
	var dolar = indo / 15000
	fmt.Printf("nilai uang dolar nya segini bos!!: %d ", dolar)

}
