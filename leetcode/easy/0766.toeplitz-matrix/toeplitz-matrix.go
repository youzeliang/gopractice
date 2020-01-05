package _766_toeplitz_matrix

func isToeplitzMatrix(matrix [][]int) bool {

	if len(matrix) == 0 {
		return false
	}

	for k, v := range matrix {
		for k1 := range v {
			if k < len(matrix)-1 && k1 < len(v)-1 {
				if matrix[k][k1] != matrix[k+1][k1+1] {
					return false
				}
			}
		}
	}
	return true
}
