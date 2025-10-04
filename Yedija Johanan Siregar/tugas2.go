package main

import "fmt"

func main() {
    const phi = 3.1415926535
    var r float64
    fmt.Printf("Masukkan nilai r bola: ")
    fmt.Scan(&r)

    volume := (4.0 / 3.0) * phi * r * r * r
    luas := 4 * phi * r * r

    // Hasil dari volume dan luas bola
    fmt.Printf("Bola dengan jejari 5 memiliki volume: %.4f, dan luas: %.4f", volume, luas)
}
