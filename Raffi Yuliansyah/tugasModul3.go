package main

import "fmt"

func main() {
	soal1()
	soal2()
	soal3()
	soal4()
	contohsoal1()
	contohsoal2()
	contohsoal3()
}

func contohsoal3() {
	var rupiah, dolar int
	fmt.Scan(&rupiah)
	dolar = rupiah / 15000
	fmt.Println(dolar)
}

func contohsoal2() {
	var alas, tinggi, luas float64
	fmt.Scan(&alas, &tinggi)
	luas = 0.5 * alas * tinggi
	fmt.Println(luas)
}

func contohsoal1() {
	var sisi, volume float64
	fmt.Scan(&sisi)
	volume = sisi * sisi * sisi
	fmt.Println(volume)
}

func soal4() {
	var celcius, reamur, fahrenheit, kelvin float64
	fmt.Print("Input, derjat temperatur Celcius = ")
	fmt.Scanln(&celcius)
	reamur = (celcius * 4) / 5
	fahrenheit = ((celcius * 9) / 5) + 32
	kelvin = celcius + 273
	fmt.Println("Output, Derajat Reamur : ", reamur)
	fmt.Println("Output, Derajat Fahrenheit : ", fahrenheit)
	fmt.Println("Output, Derajat Kelvin : ", kelvin)
}

func soal3() {
	var tahun int
	var kabisat bool
	fmt.Print("Input, cek tahun kabisat : ")
	fmt.Scanln(&tahun)
	kabisat = (tahun%400 == 0) || ((tahun%100 != 0) && (tahun%4 == 0))
	fmt.Print("Output, tahun kabisat : ", kabisat)
}

func soal2() {
	const pi = 3.1415926535
	var r, volumeBola, luasPermukaanBola float64
	fmt.Print("Input nilai r (jari-jari bola) = ")
	fmt.Scanln(&r)
	volumeBola = ((4 * pi * r * r * r) / 3)
	luasPermukaanBola = 4 * pi * r * r
	fmt.Printf("Output volume bola dengan r = %.1f adalah %.4f \n", r, volumeBola)
	fmt.Printf("Output luas permukaan bola dengan r = %.1f adalah %.4f", r, luasPermukaanBola)
}

func soal1() {
	var fx, x float64
	fmt.Print("Input nilai f(x) = ")
	fmt.Scanln(&fx)
	x = ((2 / (fx - 5)) - 5)
	fmt.Printf("Ouput nilai dari x = %.0f", x)
}
