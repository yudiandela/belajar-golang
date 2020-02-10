package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		switch i {
		case 1:
			fmt.Println("Satu")
		case 2:
			fmt.Println("Dua")
		case 3:
			fmt.Println("Tiga")
		case 4:
			fmt.Println("Empat")
		case 5:
			fmt.Println("Lima")
		default:
			fmt.Println("Lebih Dari Lima")
		}
	}
}
