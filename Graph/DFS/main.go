package main

import (
	"fmt"
)

type graph struct {
	vertices []*vertex
}
type vertex struct {
	data     int
	adjacent []*vertex
}

func main() {
	g := graph{}
	g.addVertex(1)
	g.addVertex(2)
	g.addVertex(3)
	g.addVertex(4)
	g.addVertex(5)

	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(1, 4)
	g.addEdge(5, 3)
	g.addEdge(2, 4)

	g.print()
	fmt.Println("\n DFS :")
	g.DFS(5)

}
func (g *graph) addVertex(data int) {
	if !contains(g.vertices, data) {
		g.vertices = append(g.vertices, &vertex{data: data})
	}
}
func contains(vertex []*vertex, data int) bool {
	for _, v := range vertex {
		if v.data == data {
			return true
		}
	}
	return false
}
func (g *graph) print() {
	for _, v := range g.vertices {
		fmt.Print("\n vertex", v.data, ":")
		for _, v := range v.adjacent {
			fmt.Print(" ", v.data)
		}

	}
}
func (g *graph) getVertex(data int) *vertex {
	for _, v := range g.vertices {
		if v.data == data {
			return v
		}
	}
	return nil
}
func (g *graph) addEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		fmt.Println("Error: Vertex not found.")
		return
	}

	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	toVertex.adjacent = append(toVertex.adjacent, fromVertex)
}

type Stack struct {
	arr []int
}

func (g *graph) DFS(vertex int) {
	stack := Stack{}
	visited := make(map[int]bool)
	stack.arr = append(stack.arr, vertex)
	visited[vertex] = true
	for len(stack.arr) != 0 {
		vertex = stack.arr[len(stack.arr)-1]
		stack.arr = stack.arr[:len(stack.arr)-1]
		fmt.Print(vertex, " ")
		for _, v := range g.getVertex(vertex).adjacent {
			if !visited[v.data] {
				visited[v.data] = true
				stack.arr = append(stack.arr, v.data)
			}
		}
	}
}

// Recursion
// func (g *graph) DFS(vertex int) {
// 	visited := make(map[int]bool)
// 	g.recursiveDFS(vertex, visited)
// }
// func (g *graph) recursiveDFS(vertex int, visited map[int]bool) {
// 	if visited[vertex] {
// 		return 
// 	}
// 	fmt.Print(vertex, " ")
// 	visited[vertex] = true
// 	for _, v := range g.getVertex(vertex).adjacent {
// 		g.recursiveDFS(v.data, visited)
// 	}
// }
