package main

import (
	"fmt"
	"math"
	"skillbox/21_anonymous_func/fn/fn"
)

func main() {
	//1. Расчёт по формуле (анонимная функция)
	S := func(x int16, y int8, z float32) float64 {
		return 2*float64(x) + math.Pow(float64(y), 2) - 3/float64(z)
	}(2, 3, 1.5)
	fmt.Println("S:", S)

	//2. Анонимные функции
	fn.AcceptAnonymousFunc(func(x, y int) int {
		return x + y
	})

	fn.AcceptAnonymousFunc(func(x, y int) int {
		return x * y
	})

	fn.AcceptAnonymousFunc(func(x, y int) int {
		return x - y
	})
}
