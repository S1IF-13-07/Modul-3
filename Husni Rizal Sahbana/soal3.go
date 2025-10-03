package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)

	// langsung cek logika kabisat
	hasil := tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0)
	fmt.Println(hasil)
}
