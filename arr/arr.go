package arr

//MergeSortedArrays - производит слияние двух отсортированных массивов длиной 4 и 5 в один массив длиной 9
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
