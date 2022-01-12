package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

// Define the Garden type here.

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

type Position struct {
	row int
	col int
}

type Garden map[string][]string

func (g *Garden) Plants(child string) (plants []string, ok bool) {
	plants, ok = (*g)[child]
	return
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	matrix := strings.Split(diagram, "\n")[1:]
	if len(matrix) != 2 {
		return nil, fmt.Errorf("wrong diagram format")
	}
	rowLen := len(children) * 2
	for _, row := range matrix {
		if len(row) != rowLen {
			return nil, fmt.Errorf("wrong diagram format")
		}
	}

	garden := make(Garden)
	sortedChildren := make([]string, len(children))
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)

	for i, child := range sortedChildren {
		offset := 2 * i
		if _, ok := garden[child]; ok {
			return nil, fmt.Errorf("duplicate name")
		}

		positions := []Position{
			{row: 0, col: offset},
			{row: 0, col: offset + 1},
			{row: 1, col: offset},
			{row: 1, col: offset + 1},
		}
		for _, position := range positions {
			plantName, err := getPlantName(matrix[position.row][position.col])
			if err != nil {
				return nil, err
			}
			garden[child] = append(garden[child], plantName)
		}
	}
	return &garden, nil
}

func getPlantName(code uint8) (string, error) {
	switch code {
	case 'C':
		return "clover", nil
	case 'G':
		return "grass", nil
	case 'R':
		return "radishes", nil
	case 'V':
		return "violets", nil
	default:
		return "", fmt.Errorf("wrong plant")
	}
}
