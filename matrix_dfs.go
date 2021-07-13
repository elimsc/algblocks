package main

// 二维DFS
func MatrixDFS(grid [][]byte, row, col int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return
	}
	if grid[row][col] == '2' || grid[row][col] == '0' {
		return
	}
	grid[row][col] = '2'
	MatrixDFS(grid, row-1, col)
	MatrixDFS(grid, row+1, col)
	MatrixDFS(grid, row, col-1)
	MatrixDFS(grid, row, col+1)
}
