package main

import "fmt"

func main() {
	// slice ~= array yang membedakan cuma ukuran slice dinamis
	// ketika tidak tau isi dari array yang mau di slice
    var months = [...]string {
		"Januari", 
		"Februari", 
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var sliceMonths = months[4:7]

	fmt.Println(sliceMonths) // [Mei Juni Juli]
	fmt.Println(len(sliceMonths)) // 3
	fmt.Println(cap(sliceMonths)) // 8

	// merubah index ke 5 dari array months
	months[5] = "Store Aja"
	fmt.Println(sliceMonths) // [Mei Store Aja Juli]

	sliceMonths[0] = "Merubah Mei"
	// merubah data array dari index 0 dari slice [4:7] -> 4
	fmt.Println(months) // [Januari Februari Maret April Merubah Mei Store Aja Juli Agustus September Oktober November Desember]


	var slice2  = months[10:]
	fmt.Println(slice2) // [November Desember]

	
	var slice3  = append(slice2, "Kok Gini") 
	fmt.Println(slice3) // [November Desember Kok Gini]
	slice3[1] = "Berubah"  
	fmt.Println(slice3)

	fmt.Println(slice2) // [November Desember]
	fmt.Println(months) // [Januari Februari Maret April Merubah Mei Store Aja Juli Agustus September Oktober November Desember]
    

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Ama"
	newSlice[1] = "Jani"
	fmt.Println(newSlice) // [Ama Jani]
	fmt.Println(len(newSlice)) // 2
	fmt.Println(cap(newSlice)) // 5

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice) // [Ama Jani]

	iniArray := [5]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
