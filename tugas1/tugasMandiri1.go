package main

import "fmt"

func main() {
	name := "Halo, Saf"
	age := 23
	isMuslim := true

	fmt.Printf("%s \n", name)
	fmt.Printf("%d \n", age)
	fmt.Printf("%t \n", isMuslim)

	fmt.Printf("\n")

	fmt.Printf("%T \n", name)
	fmt.Printf("%T \n", age)
	fmt.Printf("%T \n", isMuslim)
}
