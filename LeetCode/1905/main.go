package main

import "fmt"

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    m := len(grid1)
    n := len(grid1[0])
    a, b := 2, 2
    
    for i := 0; i < m; i++  {
        for j := 0; j < n; j++ {
            if grid1[i][j] == 1 {
                flood(grid1, i, j, a)
                a++
            }
            if grid2[i][j] == 1 {
                flood(grid2, i, j, b)
                b++
            }
        }
    }
    
    dict := make(map[int][][2]int)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid2[i][j] != 0 {
                dict[grid2[i][j]] = append(dict[grid2[i][j]], [2]int{i, j})
            }
        }
    }

    res := 0
    LOOP:
    for _, il := range dict {
        num := -1
        for _, g := range il {
            if grid1[g[0]][g[1]] == 0 {
                continue LOOP
            }
            if num == -1 {
                num = grid1[g[0]][g[1]]
            } else if grid1[g[0]][g[1]] != num {
                continue LOOP
            }
        }
        res++
    }
    return res
}

func flood(grid [][]int, i int, j int, n int) {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
        return
    }
    grid[i][j] = n
    flood(grid, i + 1, j, n)
    flood(grid, i - 1, j, n)
    flood(grid, i, j + 1, n)
    flood(grid, i, j - 1, n)
}

func main() {
    grid1 := [][]int{
        {1, 1, 1, 0, 0},
        {0, 1, 1, 1, 0},
        {0, 0, 0, 1, 0},
        {1, 1, 1, 0, 0},
        {1, 0, 1, 1, 1},
    }

    grid2 := [][]int{
        {1, 1, 1, 0, 0},
        {0, 1, 0, 1, 0},
        {0, 0, 0, 1, 0},
        {1, 0, 1, 0, 0},
        {1, 0, 1, 1, 1},
    }

    result := countSubIslands(grid1, grid2)
    fmt.Println("Number of sub-islands:", result)
}
