package main

import (
	"fmt"
	"math"
)

func minDist(dist [6]int, vist [6]bool) int {
	min := math.MaxInt
	var min_index int
	for v := 0; v < 6; v++ {
		if !vist[v] && dist[v] <= min {
			min = dist[v]
			min_index = v
		}
	}
	return min_index
}

func printSolution(dist [6]int) {
	fmt.Println("Vertex \t Distance from Source")
	for i := 0; i < 6; i++ {
		fmt.Println(string(i+65), "\t\t", dist[i])
	}
}

func dijikstra(g [6][6]int, src int) {
	var dist [6]int
	var vist [6]bool
	for i := 0; i < 6; i++ {
		dist[i], vist[i] = math.MaxInt, false
	}
	dist[src] = 0
	for count := 0; count < 5; count++ {
		u := minDist(dist, vist)
		vist[u] = true
		for v := 0; v < 6; v++ {
			if !vist[v] && g[u][v] != 0 && dist[u] != math.MaxInt && dist[u]+g[u][v] < dist[v] {
				dist[v] = dist[u] + g[u][v]
			}
		}
	}
	printSolution(dist)

}

func main() {
	var graph = [6][6]int{{0, 3, 5, 0, 0, 0},
		{4, 0, 11, 9, 7, 0},
		{5, 11, 0, 0, 3, 0},
		{0, 9, 0, 0, 13, 2},
		{0, 7, 3, 13, 0, 6},
		{0, 0, 0, 2, 6, 0}}
	var src int = 0
	dijikstra(graph, src)
}
