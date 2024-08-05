package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	var count int
	res := make([][]int, len(grid))
	for i := range res {
		res[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res[i][j] = 1
				fill(i, j, grid, res)
				count++
			}
		}
	}
	return count
}

func fill(sr int, sc int, grid [][]byte, res [][]int) {
	if sr < 0 || sr >= len(res) || sc < 0 || sc >= len(res[0]) {
		return
	}

	if grid[sr][sc] == '0' {
		return
	}

	res[sr][sc] = 1
	grid[sr][sc] = '0'

	fill(sr+1, sc, grid, res)
	fill(sr-1, sc, grid, res)
	fill(sr, sc+1, grid, res)
	fill(sr, sc-1, grid, res)
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	result := numIslands(grid)
	fmt.Printf("Number of islands: %d\n", result)
}
