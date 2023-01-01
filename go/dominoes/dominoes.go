package dominoes

import (
	"sort"
)

// Domino type here.
type Domino [2]uint

type Dominos []Domino

func (receiver Dominos) Clone() Dominos {
	dominos := Dominos{}

	for _, domino := range receiver {
		dominos = append(dominos, Domino{domino[0], domino[1]})
	}

	return dominos
}

func (receiver Dominos) Insert(position int, domino Domino) Dominos {
	dominos := make(Dominos, 0)

	for _, d := range receiver[:position] {
		dominos = append(dominos, d)
	}

	dominos = append(dominos, domino)

	for _, d := range receiver[position:] {
		dominos = append(dominos, d)
	}

	return dominos
}

func MakeChain(dominos Dominos) (Dominos, bool) {
	if len(dominos) == 0 {
		return nil, true
	}
	chain := createChain(dominos.Clone())
	return chain, len(chain) == len(dominos) && chain[0][0] == chain[len(chain)-1][1]
}

func createChain(dominos Dominos) Dominos {
	// Sort the dominos in ascending order based on the values on their ends
	sort.Slice(dominos, func(i, j int) bool {
		if dominos[i][0] != dominos[j][0] {
			return dominos[i][0] < dominos[j][0]
		}
		return dominos[i][1] > dominos[j][1]
	})

	// Initialize the chain with the first domino in the sorted list
	chain := Dominos{dominos[0]}
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
			if domino[0] == domino[1] {
				for j := 0; j < len(chain)-1; j++ {
					if chain[j][1] == chain[j+1][0] && chain[j][1] == domino[0] {
						chain = chain.Insert(j+1, domino)
						dominos = append(dominos[:i], dominos[i+1:]...)
						found = true
						break
					}
				}
			}
		}
		if !found {
			// No match found, the input set of dominos cannot be used to form a valid chain
			return chain
		}
	}

	return chain
}
