package main

import "fmt"

func DFS(graph map[string][]string, s, f string, visited []string) bool {
	if s == f { // проверям, равен ли начальный узел конечному
		return true
	}
	for _, val := range visited { // проверяем, посещали ли мы начальную вершину
		if s == val {
			return false
		}
	}
	visited = append(visited, s) // добавляем в список посещенных вершин
	fmt.Println(visited)

	for _, neighbor := range graph[s] { // проверям всех соседей текущей вершины
		for _, val := range visited {
			if val != neighbor { // проверяем, посещали ли мы текущую вершину
				reached := DFS(graph, neighbor, f, visited)
				if reached {
					return true
				}
			}
		}
	}

	return false

}

func main() {
	graph := make(map[string][]string)

	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	var visited []string
	fmt.Println(DFS(graph, "you", "jonny", visited))

}
