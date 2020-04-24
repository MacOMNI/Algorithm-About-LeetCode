package LeetCode

var vis [][]int
var m int
var n int

func numIslands(grid [][]byte) int {
	res := 0
	m = len(grid)
	if m <= 0 {
		return res
	}
	n = len(grid[0])
	vis = make([][]int, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			vis[i][j] = 0
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if vis[i][j] != 1 && grid[i][j] == '1' {
				res++
				dfsGrid(i, j, grid)
			}
			vis[i][j] = 1
		}
	}
	return res
}
func dfsGrid(i int, j int, grid [][]byte) {
	vis[i][j] = 1
	if j-1 >= 0 && grid[i][j-1] == '1' && vis[i][j-1] == 0 { //left
		dfsGrid(i, j-1, grid)
	}
	if j+1 < n && grid[i][j+1] == '1' && vis[i][j+1] == 0 { //right
		dfsGrid(i, j+1, grid)
	}
	if i-1 >= 0 && grid[i-1][j] == '1' && vis[i-1][j] == 0 { //up
		dfsGrid(i-1, j, grid)
	}
	if i+1 < m && grid[i+1][j] == '1' && vis[i+1][j] == 0 { //down
		dfsGrid(i+1, j, grid)
	}
}
