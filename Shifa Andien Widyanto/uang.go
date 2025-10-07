package main
import "fmt"
func main() {
	var rupiah, dolar int

	fmt.Print("Masukan nilai rupiah : ")
	fmt.Scan(&rupiah)
	dolar = rupiah / 15000
	fmt.Println(dolar)
}