package dominoes

// Domino type here.
type Domino [2]uint
type Graph map[uint][]uint

func MakeChain(dominos []Domino) ([]Domino, bool) {
	result := len(dominos) == 0 || (allEvenDegrees(dominos) && isConnected(dominos))
	return nil, result
}

func getAdjacencyList(dominos []Domino) Graph {
	graph := Graph{}
	for _, domino := range dominos {
		graph[domino[0]] = append(graph[domino[0]], domino[1])
		graph[domino[1]] = append(graph[domino[1]], domino[0])

	}
	return graph
}

func allEvenDegrees(dominos []Domino) bool {
	nodes := map[uint]uint{}
	for _, domino := range dominos {
		nodes[domino[0]]++
		nodes[domino[1]]++
	}
	for _, count := range nodes {
		if count%2 != 0 {
			return false
		}
	}
	return true
}

func isConnected(dominos []Domino) bool {
	graph := getAdjacencyList(dominos)
	visited := make([]bool, len(graph), len(graph))
	DepthFirstRecursive(graph, dominos[0][0], visited)
	for _, isVisited := range visited {
		if !isVisited {
			return false
		}
	}
	return true
}

func DepthFirstRecursive(graph Graph, node uint, visited []bool) {
	visited[node-1] = true
	for _, neighbor := range graph[node] {
		if !visited[neighbor-1] {
			DepthFirstRecursive(graph, neighbor, visited)
		}
	}
}
