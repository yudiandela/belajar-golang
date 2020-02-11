package main

import "fmt"

func main() {
	names := [4]string{
		"Yudi",
		"Marissa",
		"Yusza",
		"Razzaq",
	}

	fmt.Println(names)

	// Mengambil data array dengan index tertentu
	slice1 := names[1:3]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[1:]
	fmt.Println(slice3)
}
