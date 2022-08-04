package main

import (
	"fmt"
	"skillbox/20-2d_arrays/arr/arr"
)

func main() {
	//1. Подсчёт определителя
	m := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Determinant of a third order:", arr.ThirdOrderDeterminant(m))
	fmt.Println("************")

	//2. Умножение матриц
	x := [3][5]int{
		{-1, 4, 11, 2, 3},
		{8, 6, 0, 1, 10},
		{15, -7, 32, 5, -2},
	}
	y := [5][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}
	fmt.Println("Result matrix:")
	for _, v := range arr.Multiple(x, y) {
		fmt.Println(v)
	}
}
