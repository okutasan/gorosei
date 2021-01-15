package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 90

	var lulusanNilaiAkhir = nilaiAkhir > 80 //true
	var lulusanAbsensi = absensi > 80       //false

	var lulus = lulusanNilaiAkhir && lulusanAbsensi // true && false = false

	fmt.Println(lulus)

	if lulus == true {
		fmt.Println("Selamat anda lulus")
	}else {
		fmt.Println("Anda belom lulus")
	}
}