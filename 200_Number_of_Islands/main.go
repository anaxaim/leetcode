package main

import "fmt"

func dfs(grid [][]byte, x, y int) bool {
	if grid[x][y] != '1' {
		return false
	}
	grid[x][y] = '2'

	if x-1 >= 0 {
		dfs(grid, x-1, y)
	}

	if x+1 < len(grid) {
		dfs(grid, x+1, y)
	}

	if y-1 >= 0 {
		dfs(grid, x, y-1)
	}

	if y+1 < len(grid[x]) {
		dfs(grid, x, y+1)
	}

	return true
}

func numIslands(grid [][]byte) int {
	isl := 0

	for x := range grid {
		for y := range grid[x] {
			if dfs(grid, x, y) {
				isl++
			}
		}
	}

	return isl
}

func main() {
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	fmt.Println(numIslands(grid1)) // 1

	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Println(numIslands(grid2)) // 3
}
