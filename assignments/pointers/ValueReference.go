package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Printf("Reference of string: %v\n", &namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Printf("Pointer of string value: %v\n", (*namePointer))
	//fmt.Println(&namePointer)
	fmt.Printf("Reference of poointer string: %v\n", (&namePointer))
}
