package main

import "fmt"

func main() {
	/**
	 ** Increment Loop
	 **/
	fmt.Println("For Loop Increment")
	for i := 1; i <= 10; i++ {
		fmt.Println(i, "Hello World!")
	}

	fmt.Println("\n")

	/**
	 ** Decrement Loop
	 **/
	fmt.Println("For Loop Decrement")
	for i := 10; i >= 1; i-- {
		fmt.Println(i, "Hello World!")
	}

	fmt.Println("\n")
	fmt.Println("Loop Selesai")
}
