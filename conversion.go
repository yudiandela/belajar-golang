package main

import (
	"fmt"
	"strconv"
)

func main() {
	/**
	 ** Integers Conversion
	 **/
	var x int32 = 32
	fmt.Println(x)

	var y int64 = int64(x)
	fmt.Println(y)

	var z float32 = float32(y)
	fmt.Println(z)

	/**
	 ** String Conversion
	 ** doc https://golang.org/pkg/strconv
	 **/
	var a string = "100"
	fmt.Println("String", a) //  Menghasilkan string 100

	//(Harus Mengembalikan dua Parameter)
	// Error var b int32 = strconv.Atoi(a)
	b, _ := strconv.Atoi(a)
	fmt.Println("Integer", b) // Menghasilkan int 100

	c := strconv.Itoa(b)
	fmt.Println("String", c) // Menghasilkan string 100

	var d string = "true"
	fmt.Println("String", d) // Menghasilkan string true

	e, _ := strconv.ParseBool(d)
	fmt.Println("Boolean", e) // Menghasilkan boolean true
}
