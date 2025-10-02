package main

import (
	"fmt"
	"math"
)

func main() {
	var celsius float64

	// Meminta input suhu dalam Celcius dari pengguna
	fmt.Print("Masukkan suhu dalam Derajat Celsius: ")
	_, err := fmt.Scan(&celsius)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}

	// Menghitung konversi ke Fahrenheit dan membulatkannya ke bilangan bulat
	fahrenheitFloat := (celsius * 9.0 / 5.0) + 32
	fahrenheitInt := int(math.Round(fahrenheitFloat))

	// Menghitung konversi ke Reamur dan membulatkannya ke bilangan bulat
	reamurFloat := celsius * 4.0 / 5.0
	reamurInt := int(math.Round(reamurFloat))

	// Menghitung konversi ke Kelvin dan membulatkannya ke bilangan bulat
	kelvinFloat := celsius + 273.15
	kelvinInt := int(math.Round(kelvinFloat))

	// Menampilkan hasil konversi dalam bilangan bulat
	fmt.Printf("\nTemperatur Celsius : %.f\n", celsius)
	fmt.Printf("Derajat Fahrenheit : %d \n", fahrenheitInt)
	fmt.Printf("Derajat Reamur     : %d \n", reamurInt)
	fmt.Printf("Derajat Kelvin     : %d \n", kelvinInt)
}