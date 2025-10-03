package main

import "fmt"

func main() {
	var fx float64
	fmt.Print("Masukan bilangan rilll : ")
	fmt.Scan(&fx)
	x := (2/(fx-5) - 5)
	fmt.Printf("Nilai X nya adalah : %.0f", x)
}
