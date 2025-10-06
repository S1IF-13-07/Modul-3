package main 

import "fmt"

func main(){
	var idr float64
	var usd int
	fmt.Printf ("Masukkan jumlah uang dalam IDR: ")
	fmt.Scan(&idr)
	usd = int (idr / 15000)
	fmt.Printf ("Jumlah uang dalam USD adalah %d", usd)
}