package main

import "fmt"

const KURS_IDR_TO_USD = 15000.0

func main() {
    var idr int
    fmt.Print("Masukkan jumlah uang dalam IDR (Rupiah): ")
    
    fmt.Scan(&idr) 
    
    usd := float64(idr) / KURS_IDR_TO_USD
    fmt.Printf("Keluaran: %.0f USD\n", usd)
}