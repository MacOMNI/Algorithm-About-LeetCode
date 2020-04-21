package LeetCode

func generateMatrix(n int) [][]int {
	var matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	sum := 1
	left := 0
	right := n - 1
	up := 0
	down := n - 1
	// if n == 1 {
	// 	matrix[0][0] = 1
	// 	return matrix
	// }
	for sum <= n*n {
		for j := left; j <= right; j++ { //右
			matrix[up][j] = sum
			sum++
		}
		up++
		for j := up; j <= down; j++ { //下
			matrix[j][right] = sum

			sum++
		}
		right--
		for j := right; j >= left; j-- { //左
			matrix[down][j] = sum
			sum++
		}
		down--
		for j := down; j >= up; j-- { //上
			matrix[j][left] = sum
			sum++
		}
		left++
	}
	return matrix
}
