package main

import "fmt"

type Address struct {
   City, Province, Country string
}
// change to pointer
func ChangeCountryToIndonesia(address *Address) {
	 address.Country = "Indonesia"
}

func main() {
   address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
   address2 := &address1 //pointer &
   address3 := &address1 //pointer

   address2.City = "Bandung"

   // address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}
   *address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

   fmt.Println(address1)
   fmt.Println(address2)
   fmt.Println(address3)

   var address4 *Address = new(Address)
   address4.City = "Jakarta"
   fmt.Println(address4)

   var alamat = Address{
	  City: "Bekasi",
	  Province: "Jawa Barat",
	  Country: "",
   }
   //add pointer to alamat
   ChangeCountryToIndonesia(&alamat)
   fmt.Println(alamat)
}