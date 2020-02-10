package main

import (
	"fmt"
)

func main() {
	/**
	 ** Logic If
	 **
	 ** Studi Kasus
	 ** jika i % 3 && i % 2 print FooBar
	 ** jika i % 3 print Bar
	 ** jika i % 2 print Foo
	 **/
	for i := 1; i <= 30; i++ {
		if i%3 == 0 && i%2 == 0 {
			fmt.Println(i, "FooBar")

		} else if i%3 == 0 {
			fmt.Println(i, "Bar")

		} else if i%2 == 0 {
			fmt.Println(i, "Foo")

		} else {
			fmt.Println(i)
		}
	}
}
