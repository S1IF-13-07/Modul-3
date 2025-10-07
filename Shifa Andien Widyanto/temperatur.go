package main
import "fmt"
func main(){
	var c float64

	fmt.Print("Temperatur Celcius : ")
	fmt.Scan(&c)

	var reamur = c * 4.0/5.0
	var fahrenheit = c * 9.0/5.0 + 32.0
	var kelvin = c + 273.0

	fmt.Printf("Derajat Reamur : %.0f\n ", reamur)
	fmt.Printf("Derajat Fahreinheit : %.0f\n ", fahrenheit)
	fmt.Printf("Derajat Kelvin : %.0f\n ", kelvin)
}