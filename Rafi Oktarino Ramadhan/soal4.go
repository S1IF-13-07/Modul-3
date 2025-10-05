package main
import "fmt"

func main(){
	var celcius, fahrenheit, reamur, kelvin float64

	fmt.Print("Temperatur Celcius : ")
	fmt.Scan(&celcius)

	reamur = celcius * 4 / 5
    fahrenheit = (celcius * 9 / 5) + 32
    kelvin = celcius + 273

	fmt.Printf("Derajat Reamur = %.0f\n", reamur)
    fmt.Printf("Derajat Fahrenheit = %.0f\n", fahrenheit)
    fmt.Printf("Derajat Kelvin = %.0f\n", kelvin)

}
