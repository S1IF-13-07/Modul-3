package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Masukkan tahun :")
	fmt.Scanln(&tahun)

	var kabisat = (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)
	fmt.Printf("apakah %d tahun kabisat ? %t\n", tahun, kabisat)
}
