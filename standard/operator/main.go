package main

import "fmt"

func main() {
	var x uint8 = 200
	fmt.Printf("x: %v\n", x)
	fmt.Printf("x(binary): %08b\n", x)

	y := ^x
	fmt.Printf("y: %v\n", y)
	fmt.Printf("y(binary): %08b\n", y)
}