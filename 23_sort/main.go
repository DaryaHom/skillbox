package main

import (
	"fmt"
	"skillbox/23_sort/sort"
)

func main() {
	//1. Чётные и нечётные
	array := []int{1, 8, 12, 4, 6, 11, 9, 0, 3, 8, 7, 10}
	even, odd := sort.Separate(array)
	fmt.Println("Even numbers:", even)
	fmt.Println("Odd numbers:", odd)

	fmt.Println("************")

	//2. Поиск символов в нескольких строках
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := []rune{'W', 'E', 'L', 'и', 'М'}
	array2d, err := sort.ParseTest(sentences, chars)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for i, r := range chars {
		fmt.Printf("%q position %v\n", r, array2d[i])
	}
}
