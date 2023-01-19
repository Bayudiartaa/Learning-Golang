package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var kkm = 80

	var lulusKkm bool = nilaiAkhir > 80
	var lulusNilaiAkhir bool = kkm > 70	

	var lulus bool = lulusKkm && lulusNilaiAkhir
	fmt.Println(lulus) // true

	// shorthand code
	fmt.Println(nilaiAkhir >= 80 && kkm >= 80) //true
}