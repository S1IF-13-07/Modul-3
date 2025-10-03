package main
import "fmt"

func main(){
	var rupiah int
	fmt.Print("Masukan Nilai Uang Dalam Rupiah: ")
	fmt.Scan(&rupiah)
	var dolar = rupiah / 15000
	fmt.Print("Maka Nilai Uang Dalam Dolar: ", dolar)
}