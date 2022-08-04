package main

import (
	"fmt"
	"skillbox/22_search/sch"
)

func main() {
	//1. Подсчёт чётных и нечётных чисел в массиве
	array := sch.CreateUnsortedArray()
	fmt.Println("Array:", array)
	number := sch.GetCustomerNumber()
	sch.NumbersAfter(array, number)

	fmt.Println("************************")

	//2. Нахождение первого вхождения числа в упорядоченном массиве
	array = sch.CreateSortedArray()
	//Такой массив не может содержать дубликатов
	//Массив для проверки алгоритма поиска индекса первого вхождения:
	//array = [12]int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Array:", array)
	number = sch.GetCustomerNumber()
	sch.FirstOccurrence(array, number)
}
