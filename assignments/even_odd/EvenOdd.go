package main

import (
	"fmt"
)

func main() {
	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range num {
		printNum(n)
	}

	fmt.Println("")

	numLoop := []int{}
	for i := 0; i < 10; i++ {
		numLoop = append(numLoop, i)
	}

	for i := 0; i < len(numLoop); i++ {
		printNum(numLoop[i])
	}
}

func printNum(n int) {
	if (n % 2) == 0 {
		fmt.Printf("%v is even\n", n)
	} else {
		fmt.Printf("%v is odd\n", n)
	}
}
