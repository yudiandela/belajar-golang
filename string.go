package main

import "fmt"

func main() {
	/**
	 ** Manipulasi String di Go Lang
	 **/
	fmt.Println("Belajar Go-Lang")                           // Menampilkan String
	fmt.Println("Belajar" + " " + "Pemograman" + " " + "GO") // String Concatenate
	fmt.Println(len("Belajar Pemograman GO"))                // Menampilkan Panjang String

	/**
	 ** Mengambil Index Tertentu
	 **/
	fmt.Println("Belajar Pemograman GO"[0])   // @return byte
	fmt.Println("Belajar Pemograman GO"[0:5]) // @return Belaj
	fmt.Println("Belajar Pemograman GO"[5:])  // @return ar Pemograman GO
	fmt.Println("Belajar Pemograman GO"[:10]) // @return Belajar Pe
}
