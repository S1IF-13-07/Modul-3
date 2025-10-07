package main

import "fmt"


func main (){
	var sisi, volume float64

	fmt.Print("Masukkan sisi kubus: ")
	fmt.Scan(&sisi)
	volume = (sisi * sisi * sisi)

	fmt.Printf("volume kubus = %.1f\n", volume)

}
 
