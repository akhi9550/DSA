package main

import (
	"fmt"
)

func regionsBySlashes(grid []string) int {
	n := len(grid)
	a := make([][]int, 3*n)
	for i := 0; i < len(a); i++ {
		a[i] = make([]int, 3*n)
	}
	fill(grid, a)
	return count(a)
}

func fill(grid []string, a [][]int) {
	n := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == ' ' {
				continue
			} else if grid[i][j] == '/' {
				a[3*i+2][3*j+0] = 1
				a[3*i+1][3*j+1] = 1
				a[3*i+0][3*j+2] = 1
			} else if grid[i][j] == '\\' {
				a[3*i+0][3*j+0] = 1
				a[3*i+1][3*j+1] = 1
				a[3*i+2][3*j+2] = 1
			} else {
				panic("unknown char")
			}
		}
	}
}

func count(a [][]int) int {
	m := len(a)
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == 0 {
				flood(a, i, j)
				res++
			}
		}
	}
	return res
}

func flood(a [][]int, i, j int) {
	m := len(a)
	if i < 0 || j < 0 || i >= m || j >= m {
		return
	}
	if a[i][j] != 0 {
		return
	}
	a[i][j] = 2
	flood(a, i, j+1)
	flood(a, i, j-1)
	flood(a, i-1, j)
	flood(a, i+1, j)
}

func main() {
	grid := []string{
		" /",
		"/ ",
	}
	fmt.Println("Number of regions:", regionsBySlashes(grid))
}
