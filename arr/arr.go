package arr

//ThirdOrderDeterminant - вычисляет определитель матрицы размером 3×3
func ThirdOrderDeterminant(matrix [3][3]int) int {
	a := matrix[0][0] * matrix[1][1] * matrix[2][2]
	b := matrix[0][0] * matrix[1][2] * matrix[2][1]
	c := matrix[0][1] * matrix[1][0] * matrix[2][2]
	d := matrix[0][1] * matrix[1][2] * matrix[2][0]
	e := matrix[0][2] * matrix[1][0] * matrix[2][1]
	f := matrix[0][2] * matrix[1][1] * matrix[2][0]
	result := a - b - c + d + e - f
	return result
}

// Multiple - умножает две матрицы размерами 3×5 и 5×4
func Multiple(m1 [3][5]int, m2 [5][4]int) (result [3][4]int) {
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m2[j]); j++ {
			for k := 0; k < len(m2); k++ {
				result[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return
}

//MergeSortedArrays - производит слияние двух отсортированных массивов длиной 4 и 5 в один массив длиной 9
//Код для проверки:
//a := [...]int{-8, 5, 5, 8}
//b := [...]int{0, 4, 5, 6, 9}
//fmt.Println("First sorted array", a)
//fmt.Println("Second sorted array", b)
//fmt.Println("Merged array:", arr.MergeSortedArrays(a, b))
func MergeSortedArrays(firstArray [4]int, secondArray [5]int) (mergedArray [9]int) {
	i, j, m := 0, 0, 0
	for i < len(firstArray) && j < len(secondArray) {
		if firstArray[i] < secondArray[j] {
			mergedArray[m] = firstArray[i]
			i++
		} else if firstArray[i] > secondArray[j] {
			mergedArray[m] = secondArray[j]
			j++
		} else {
			mergedArray[m] = firstArray[i]
			i++
			m++
			mergedArray[m] = secondArray[j]
			j++
		}
		m++
	}
	for i < len(firstArray) {
		mergedArray[m] = firstArray[i]
		i++
		m++
	}
	for j < len(secondArray) {
		mergedArray[m] = secondArray[j]
		j++
		m++
	}
	return
}

//BubbleSort - сортирует массив длиной 6 пузырьком.
//Код для проверки:
//array := [...]int{6, 6, 0, 2, 6, -11}
//fmt.Println("Unsorted array", array)
//fmt.Println("Sorted array:", arr.BubbleSort(array))
func BubbleSort(array [6]int) [6]int {
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
	return array
}
