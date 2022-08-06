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
	array = [10]int{10, 2, 10, 3, 1, 2, 5, 8, 4}
	fmt.Println("Unsorted array:", array)
	f := func(array [10]int) [10]int {
		var isArraySorted bool
		for !isArraySorted {
			isArraySorted = true
			for i := 1; i < len(array); i++ {
				if array[i-1] > array[i] {
					array[i-1], array[i] = array[i], array[i-1]
					isArraySorted = false
				}
			}
		}
		for i := 0; i < len(array)/2; i++ {
			array[i], array[len(array)-1-i] = array[len(array)-1-i], array[i]
		}
		return array
	}
	fmt.Println("Bubble sorted & reversed array:", f(array))
}
