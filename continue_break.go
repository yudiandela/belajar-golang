package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		// Menghasilkan angka ganjil
		// menggunakan fungsi continue
		if i%2 == 0 {
			continue
		}

		// Menghentikan Looping
		// menggunakan break
		if i == 51 {
			break
		}

		fmt.Println(i)
	}
}
