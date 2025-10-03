package main 
import "fmt"

func main() {
	var sisi int
	fmt.Print("Masukkan nilai sisi: ")
	fmt.Scan(&sisi)
	
	volumeKubus := (sisi * sisi * sisi) 
	fmt.Println("Volume kubusnya adalah:", volumeKubus)
}