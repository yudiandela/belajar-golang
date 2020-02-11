package main

import "fmt"

func main() {
	// var names [4]string

	// names[0] = "Yudi Andela"
	// names[1] = "Marissa Masir"
	// names[2] = "Muhammad Yusza Mahardika"
	// names[3] = "Muhammad Razzaq"

	names := [4]string{
		"Yudi Andela",
		"Marissa Masir",
		"Muhammad Yusza Mahardika",
		"Muhammad Razzaq",
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	fmt.Println("\n")

	/**
	 * sama dengan Foreach
	 **/
	for i, val := range names {
		fmt.Println("Index", i, "=", val)
	}

	fmt.Println("\n")

	// Tanpa menggunakan index
	for _, val := range names {
		fmt.Println(val)
	}

}
