package main

import (
	"fmt"
	"skillbox/19_simple_algo/alg/alg"
)

func main() {
	//1. Слияние отсортированных массивов
	a := [...]int{-8, 5, 5, 8}
	b := [...]int{0, 4, 5, 6, 9}
	fmt.Println("First sorted array", a)
	fmt.Println("Second sorted array", b)
	fmt.Println("Merged array:", alg.MergeSortedArrays(a, b))

	fmt.Println("************")

	//2. Сортировка пузырьком
	array := [...]int{6, 6, 0, 2, 6, -11}
	fmt.Println("Unsorted array", array)
	fmt.Println("Sorted array:", alg.BubbleSort(array))
}
