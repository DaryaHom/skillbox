package main

import (
	"fmt"
	"skillbox/24_slices/sl"
)

func main() {
	//1. Сортировка вставками
	array := [10]int{10, 2, 10, 3, 1, 2, 5, 8, 4}
	fmt.Println("Unsorted array:", array)
	fmt.Println("Sorted by insertion array:", sl.InsertionSort(array))

	fmt.Println("**********")
	//2. Анонимные функции
	f := func(slice ...int) []int {
		fmt.Println("Unsorted array:", slice)
		var isArraySorted bool
		for !isArraySorted {
			isArraySorted = true
			for i := 1; i < len(slice); i++ {
				if slice[i-1] > slice[i] {
					slice[i-1], slice[i] = slice[i], slice[i-1]
					isArraySorted = false
				}
			}
		}
		for i := 0; i < len(slice)/2; i++ {
			slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
		}
		return slice
	}(10, 2, 10, 3, 1, 2, 5, 8, 4)
	fmt.Println("Bubble sorted & reversed array:", f)
}
