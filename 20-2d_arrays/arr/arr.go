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
