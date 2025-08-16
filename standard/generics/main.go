package main

import "fmt"

type Celsius float64
type Meter int64

type NumType interface {
	~int | ~int64 | ~float32 | ~float64
}

func Min[T NumType](a, b T) T {
	if a < b { return a }
	return b
}

func main() {
	var a Celsius = 30.8
	var b Celsius = 29.4
	fmt.Printf("Celsius Min: %f\n", Min(a, b))

	var x Meter = 89
	var y Meter = 193
	fmt.Printf("Meter Min: %d\n", Min(x, y))
}