package main

import "fmt"

func main() {
    var celsius float64

    // Input suhu dalam Celsius
    fmt.Print("Temperatur Celsius: ")
    fmt.Scan(&celsius)

    // Konversi ke Reamur, Fahrenheit, dan Kelvin
    reamur := celsius * 4.0 / 5.0
    fahrenheit := (celsius * 9.0 / 5.0) + 32
    kelvin := celsius + 273

    // Tampilkan hasil
    fmt.Printf("Derajat Reamur: %.2f\n", reamur)
    fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
    fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}
