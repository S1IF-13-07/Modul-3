package main
import "fmt"

func main() {
    var sisi, volume float64

    fmt.Print("Masukan panjang sisi kubus: ")
    fmt.Scan(&sisi)
    volume = sisi * sisi * sisi
    fmt.Println("Volume kubus adalah ", volume)
}
