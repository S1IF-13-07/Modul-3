package main
import"fmt"
func main(){
	var rp float64
	const kurs = 15000
	fmt.Printf("Masukkan jumlah uang (IDR): ")
	fmt.Scanln(&rp)
	usd := rp/kurs
	fmt.Printf("Hasil konversi: %.0f USD\n", usd)
}