package main 

import "fmt"

func main() {
	var IDR int 
	var USD int = 15000
	fmt.Print("Masukkan mata uang IDR : ")
	fmt.Scan(&IDR)
	konvers := IDR / USD
	fmt.Printf("Konversi dari mata uang IDR ke USD adalah %d", konvers)
}