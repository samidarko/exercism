package dominoes

import "sort"

// Domino type here.
type Domino [2]uint
type Graph map[uint][]uint

func MakeChain(dominos []Domino) ([]Domino, bool) {
	if len(dominos) == 0 {
		return nil, true
	}
	isAllEvenDegrees := allEvenDegrees(dominos[:])
	if !isAllEvenDegrees {
		return nil, false
	}
	if isAllEvenDegrees && isConnected(dominos[:]) {
		return createChain(dominos[:]), true
		//return nil, true
	} else {
		return nil, false
	}
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

func createChain(dominos []Domino) []Domino {
	// Sort the dominos in ascending order based on the values on their ends
	sort.Slice(dominos, func(i, j int) bool {
		if dominos[i][0] != dominos[j][0] {
			return dominos[i][0] < dominos[j][0]
		}
		return dominos[i][1] < dominos[j][1]
	})

	// Initialize the chain with the first domino in the sorted list
	chain := []Domino{dominos[0]}
	dominos = dominos[1:]

	for len(dominos) > 0 {
		// Try to find a matching domino for one of the ends in the chain
		found := false
		for i, domino := range dominos {
			lastDots := chain[len(chain)-1][1]
			if lastDots == domino[0] {
				// Match found, add the new domino to the chain and remove it from the list
				chain = append(chain, domino)
				dominos = append(dominos[:i], dominos[i+1:]...)
				found = true
				break
			}
			if lastDots == domino[1] {
				// Match found, flip the domino and add it to the chain
				domino[0], domino[1] = domino[1], domino[0]
				chain = append(chain, domino)
				dominos = append(dominos[:i], dominos[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			// No match found, the input set of dominos cannot be used to form a valid chain
			return nil
		}
	}

	return chain
}
