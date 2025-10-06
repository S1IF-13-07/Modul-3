package main

import "fmt"

func main() {
	var idr int
	var usd int

	fmt.Print("masukkan nilai uang dalam IDR = Rp ")
	fmt.Scan(&idr)

	usd = int (idr / 15000)
	fmt.Printf("nilai uang idr dalam usd adalah = %d", usd)
}