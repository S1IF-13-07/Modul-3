package main
import "fmt"

func main() {
 var rupiah, dolar int

fmt.Print("Masukan uang dalam satuan IDR: ")
 fmt.Scan(&rupiah)
 dolar = rupiah / 15000
 fmt.Println("Hasil uang dalam satuan USD adalah ", dolar)
}
