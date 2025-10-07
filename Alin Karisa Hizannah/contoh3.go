package main

import "fmt"

func main() {
	//IDR - USD kurs = 15.000 IDR/USD
	var (
		idr, usd int
	)
	//INPUT TERDIRI DARI BILANGAN BULAT MENYATAKAN UANG IDR
	fmt.Print("Masukkan uang (IDR) = ")
	fmt.Scan(&idr)
	usd = idr / 15000
	//OUTPUT TERDIRI DARI BILANGAN YANG MENYATAKAN UANG DALAM SATUAN USD
	fmt.Print("Hasil Kurs (USD)= ", usd)

}
