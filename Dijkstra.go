package main

import (
	"fmt"
	"math"
)


func Dijkstra(costs map[string]int, graph map[string]map[string]int,
	parents map[string]string) (map[string]int, map[string]string) {

	processed := make(map[string]bool)

	node := find_lowest_cost_node(costs, processed)
	for node!=""{
		cost:=costs[node]
		neighbors:=graph[node]
		for n:=range neighbors{
			new_cost := cost+neighbors[n]
			if costs[n] > new_cost{
				costs[n] = new_cost
				parents[n] = node
			}
		}
		processed[node] = true
		node = find_lowest_cost_node(costs, processed)
	}
	return costs,parents

}

func find_lowest_cost_node(costs map[string]int, processed map[string]bool) string{
	lowest_cost:= math.MaxInt64
	lowest_cost_node:=""
	for node,cost:=range costs{
		if _, inProcessed := processed[node]; cost < lowest_cost && !inProcessed {
			lowest_cost = cost
			lowest_cost_node = node
		}
	}
	return lowest_cost_node
}
func main() {


	graph := make(map[string]map[string]int)
	graph["start"] = map[string]int{}
	graph["start"]["a"] = 5
	graph["start"]["b"] = 2

	graph["a"] = map[string]int{}
	graph["a"]["d"] = 4
	graph["a"]["c"] = 2


	graph["b"] = map[string]int{}
	graph["b"]["c"] = 7
	graph["b"]["a"] = 8

	graph["c"] = map[string]int{}
	graph["c"]["finish"] = 1

	graph["d"] = map[string]int{}
	graph["d"]["finish"] = 3
	graph["d"]["c"] = 6

	graph["finish"] = map[string]int{}

	costs:=map[string]int{
		"a" : 5,
		"b" : 2,
		"c" : math.MaxInt64,
		"d" : math.MaxInt64,
		"finish" : math.MaxInt64,
	}

	parents:=map[string]string{
		"a" : "start",
		"b" : "start ",
		"c" : "",
		"finish": "",

	}
	Dijkstra(costs,graph,parents)
	fmt.Println(costs, parents)
}
