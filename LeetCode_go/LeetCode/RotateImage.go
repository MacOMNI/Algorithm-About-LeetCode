package LeetCode

func rotate(matrix [][]int) {
	n := len(matrix)
	//len := len(matrix) - 1

	for i := 0; i < n-1; i++ {
		for j := i; j < n-i-1; j++ {
			ttmp := matrix[i][j]
			rtmp := matrix[j][n-i-1]
			btmp := matrix[n-i-1][n-1-j]
			ltmp := matrix[n-1-j][i]
			matrix[j][n-i-1] = ttmp
			matrix[n-i-1][n-1-j] = rtmp
			matrix[n-1-j][i] = btmp
			matrix[i][j] = ltmp
		}

	}
}
