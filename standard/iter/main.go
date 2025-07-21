package main

import "fmt"

func Fibonacci(limit int) func(yield func(int) bool) {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for a <= limit {
			fmt.Printf("イテレータ: %d をyield\n", a)
			// yieldがfalseを返したら、即座にreturnする
			if !yield(a) {
				fmt.Println("イテレータ: yieldがfalseを返したので停止")
				return
			}
			a, b = b, a+b
		}
	}
}


func main() {
	for num := range Fibonacci(100) {
		fmt.Printf("ループ: %d を受け取りました\n", num)
		// 受け取った値が20を超えたらループを抜ける
		if num > 20 {
			fmt.Println("ループ: 20を超えたのでbreak")
			break
		}
	}
}