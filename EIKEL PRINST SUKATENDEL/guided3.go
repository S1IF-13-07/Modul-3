	package main

import "fmt"

func main() {
	const kurs = 15000
	var rp int
	fmt.Print("Masukkan nilai Rupiah: ")
	fmt.Scan(&rp)
	fmt.Printf("Hasil konversi: %d USD\n", rp/kurs)
}