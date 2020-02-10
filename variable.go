package main

import "fmt"

func main() {
	/**
	 ** Mendefinisikan Variable
	 **/
	var x string = "Hello World!"
	fmt.Println(x)

	// Ubah data variable
	x = "Halo Indonesia!"
	fmt.Println(x)

	var nama string
	var usia int
	nama = "Yudi"
	usia = 30

	fmt.Println("Nama", nama, "usia saya", usia, "thn")

	namaLengkap := "Yudi Andela"
	umur := 30
	nikah := true
	fmt.Println("Nama", namaLengkap, "umur", umur, "menikah?", nikah)
}
