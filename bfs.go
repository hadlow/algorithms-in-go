package main

import "fmt"

func bfs(n int, g [][]int) {
	visited := make([]int, n)

	// Add start node to queue
	q := NewQueue()
	q.Enqueue(0)

	visited[0] = 1
	var node int

	for !q.IsEmpty() {
		node = q.Peek()
		q.Dequeue()
		fmt.Println(node)

		for _, x := range visited {
			if g[node][x] == 1 && visited[x] == 0 {
				visited[x] = 1

				q.Enqueue(x)
			}
		}
	}
}

func main() {
	n := 4

	g := [][]int{
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{1, 0, 0, 1},
		{0, 0, 0, 1},
	}

	bfs(n, g)
}
