package main

import "fmt"

func main (){
	var idr, usd, kurs int

	fmt.Print("Masukkan jumlah uang dalan IDR: ")
	fmt.Scan(&idr)
	fmt.Print("Masukkan nilai kurs: ")
	fmt.Scan(&kurs)
	usd = idr / kurs
	fmt.Printf("Hasil : %d\n", usd)
}