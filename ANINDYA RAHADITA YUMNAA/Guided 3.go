package main

import "fmt"

func main() {
	var idr, usd, kurs int
	fmt.Print("Masukkan jumlah uang dalam IDR: ")
	fmt.Scan(&idr)
	fmt.Print("Masukkan nilai tukar USD ke IDR: ")
	fmt.Scan(&kurs)
	usd = idr / kurs
	fmt.Println("Jumlah uang dalam USD adalah:", usd)
}