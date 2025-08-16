package main

import "fmt"

func main() {
	Sample2()
}

func Sample1() {
	var x uint8 = 200
	fmt.Printf("x: %v\n", x)
	fmt.Printf("x(binary): %08b\n", x)

	y := ^x // ビット反転
	fmt.Printf("y: %v\n", y)
	fmt.Printf("y(binary): %08b\n", y)

	var z int = 200
	fmt.Printf("z: %v\n", z)
	fmt.Printf("z(binary): %08b\n", z)

	w := ^z + 1 // 二の補数
	fmt.Printf("w: %v\n", w)
	fmt.Printf("w(binary): %08b\n", w)
}

func Sample2() {
	var x uint8 = 200
	fmt.Printf("x: %v\n", x)
	fmt.Printf("x(binary): %08b\n", x)

	var y uint8 = 100
	fmt.Printf("y: %v\n", y)
	fmt.Printf("y(binary): %08b\n", y)

	z := x &^ y // ビットAND NOT
	fmt.Printf("z: %v\n", z)
	fmt.Printf("z(binary): %08b\n", z)
}
