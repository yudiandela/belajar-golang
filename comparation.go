package main

import "fmt"

func main() {
	a := 1
	b := 2

	ab := a == b    // sama dengan
	fmt.Println(ab) // @return false

	c := 1
	d := 2

	cd := c != d    // tidak sama dengan
	fmt.Println(cd) // @return true

	e := 1
	f := 2

	ef := e < f     // lebih kecil dari
	fmt.Println(ef) // @return true

	g := 1
	h := 2

	gh := g > h     // lebih besar dari
	fmt.Println(gh) // @return false

	i := 1
	j := 2

	ij := i <= j    // lebih kecil atau sama dengan dari
	fmt.Println(ij) // @return true

	k := 1
	l := 2

	kl := k >= l    // lebih besar atau sama dengan dari
	fmt.Println(kl) // @return false
}
