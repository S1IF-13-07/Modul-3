package main
import "fmt"
func main() {
	var sisi int
	var volume float64

	fmt.Print("Masukan sisi kubus : ")
	fmt.Scan(&sisi)

	volume = (float64(sisi) * float64(sisi) * float64(sisi)) + 0.5
	fmt.Printf("Volume kubus = %.1f\n", volume)
}