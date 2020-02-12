package main

import "fmt"

func main() {
	/**
	 * Mendeklarasikan Slice
	 **/
	// Cara Pertama
	sliceDefine1 := []string{
		"Yudi",
		"Marissa",
		"Yusza",
		"Razzaq",
	}

	fmt.Println(sliceDefine1)
	// @return [Yudi Marissa Yusza Razzaq]

	// Cara Kedua
	sliceDefine2 := make([]string, 4)
	sliceDefine2[0] = "Yudi"
	sliceDefine2[1] = "Marissa"
	sliceDefine2[2] = "Yusza"
	sliceDefine2[3] = "Razzaq"

	fmt.Println(sliceDefine2)
	// @return [Yudi Marissa Yusza Razzaq]

	/**
	 * Manipulasi Data
	 **/
	// Mengubah Seluruh data
	sliceDefine2[0] = "Ody"
	sliceChange2 := sliceDefine2

	fmt.Println(sliceDefine2)
	// @return [Ody Marissa Yusza Razzaq]
	fmt.Println(sliceChange2)
	// @return [Ody Marissa Yusza Razzaq]

	// Menambahkan data dengan method append()
	sliceChange3 := append(sliceDefine1, "Andela", "Masir")

	fmt.Println(sliceDefine1)
	// @return [Yudi Marissa Yusza Razzaq]
	fmt.Println(sliceChange3)
	// @return [Yudi Marissa Yusza Razzaq Andela Masir]

	// Copy data dengan method copy()
	sliceChange4 := make([]string, 4)
	// Copy data dari sliceDefine1 ke sliceChange4
	copy(sliceChange4, sliceDefine1)

	fmt.Println(sliceChange4)
	// @return [Yudi Marissa Yusza Razzaq]

	// Merubah data sliceDefine[0]
	// Setelah di copy
	sliceDefine1[0] = "Andela"

	fmt.Println(sliceDefine1)
	// @return [Andela Marissa Yusza Razzaq]

	// Tidak ada perubahan
	fmt.Println(sliceChange4)
	// @return [Yudi Marissa Yusza Razzaq]
}
