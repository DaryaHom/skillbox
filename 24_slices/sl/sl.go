package sl

const n = 10

//InsertionSort - sorts an array by insertions
func InsertionSort(array [n]int) [n]int {
	for i := 0; i < n; i++ {
		value := array[i]
		j := i - 1
		for ; j >= 0; j-- {
			if value < array[j] {
				array[j+1] = array[j]
			} else {
				break
			}
		}
		array[j+1] = value
	}
	return array
}
