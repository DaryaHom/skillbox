package sch

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 12

//CreateUnsortedArray - создаёт массив из 12 неупорядоченных чисел
func CreateUnsortedArray() [n]int {
	var array [n]int
	rand.Seed(time.Now().UnixNano())
	for i := range array {
		array[i] = rand.Intn(n) + 1
	}
	return array
}

//CreateSortedArray - создаёт массив из 12 упорядоченных чисел
func CreateSortedArray() [n]int {
	var array [n]int
	rand.Seed(time.Now().UnixNano())
	for i := range array {
		array[i] = n*i + rand.Intn(n)
	}
	return array
}

//GetCustomerNumber - получает число от пользователя
func GetCustomerNumber() int {
	var number int
	fmt.Println("Enter an integer please:")
	fmt.Scan(&number)
	return number
}

//NumbersAfter - выводит количество чисел в массиве, после заданного
func NumbersAfter(array [n]int, number int) {
	index := 0
	for i, v := range array {
		if number == v {
			index = i
			fmt.Println("Index of your number:", index)
			fmt.Printf("After your number, there are %d elements in the array\n", n-1-i)
			return
		}
	}
	fmt.Println("Number not found:", index)
}

//FirstOccurrence - выводит индекс первого вхождения числа в массив
func FirstOccurrence(array [n]int, number int) {
	index := -1
	min := 0
	max := n - 1
	for max >= min {
		middle := (max + min) / 2
		if array[middle] == number {
			index = middle
			for i := middle; i > 0; i-- {
				if array[i] == array[i-1] {
					index = i - 1
				} else {
					break
				}
			}
			break
		} else if array[middle] > number {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}

	if index >= 0 {
		fmt.Println("First occurrence index:", index)
	} else {
		fmt.Println("Number not found:", index)
	}
}
