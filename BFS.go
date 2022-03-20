package main

import "fmt"

func BFS(graph map[string][]string, s, f string) bool {
	var visited []string
	var queue []string
	queue = append(queue, graph[s]...)

	for len(queue) > 0 {
		person := queue[0]
		if person == f {
			return true
		}
		queue = append(queue, graph[person]...) // добавляем всех соседей текущего узла
		visited = append(visited, person)       // добавляем уже проверенные узлы
		queue = queue[1:]                       // извлекаем из очереди

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

	fmt.Println(BFS(graph, "bob", "jonny"))

}
