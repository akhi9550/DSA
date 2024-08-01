package main

import (
	"fmt"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	q := []*Node{}
	mp := map[*Node]*Node{}

	mp[node] = &Node{Val: node.Val}

	copyNeighbors(node, mp, &q)

	for len(q) != 0 {
		front := q[len(q)-1]
		q = q[:len(q)-1]

		copyNeighbors(front, mp, &q)
	}

	return mp[node]
}

func copyNeighbors(node *Node, mp map[*Node]*Node, q *[]*Node) {
	for _, n := range node.Neighbors {
		if _, ok := mp[n]; !ok {
			mp[n] = &Node{Val: n.Val}
			*q = append(*q, n)
		}
		mp[node].Neighbors = append(mp[node].Neighbors, mp[n])
	}
}

func printGraph(node *Node) {
	visited := map[*Node]bool{}
	var dfs func(*Node)
	dfs = func(n *Node) {
		if n == nil || visited[n] {
			return
		}
		visited[n] = true
		fmt.Printf("Node %d: ", n.Val)
		for _, neighbor := range n.Neighbors {
			fmt.Printf("%d ", neighbor.Val)
		}
		fmt.Println()
		for _, neighbor := range n.Neighbors {
			dfs(neighbor)
		}
	}
	dfs(node)
}

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	fmt.Println("Original graph:")
	printGraph(node1)

	clonedGraph := cloneGraph(node1)

	fmt.Println("\nCloned graph:")
	printGraph(clonedGraph)
}
