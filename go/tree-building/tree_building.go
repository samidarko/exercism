package tree

import (
	"fmt"
	"sort"
)

// Record type
type Record struct {
	ID     int
	Parent int
}

// Node type
type Node struct {
	ID       int
	Children []*Node
}

// Append a record to a node
func (n *Node) Append(r *Record) error {
	if r.ID == r.Parent || r.ID < r.Parent {
		return fmt.Errorf("invalid node")
	}

	if r.Parent == n.ID {
		node := &Node{ID: r.ID}
		if n.IsChild(node) {
			return fmt.Errorf("duplicate node")
		}
		n.Children = append(n.Children, node)
		return nil
	}

	for _, c := range n.Children {
		if err := c.Append(r); err != nil {
			return err
		}
	}

	return nil
}

// IsChild return true if Node has target Node as child
func (n *Node) IsChild(c *Node) bool {
	for _, child := range n.Children {
		if child.ID == c.ID {
			return true
		}
	}

	return false
}

// Build a Node from a list of Record
func Build(records []Record) (*Node, error) {
	sort.Slice(records, func(i, j int) bool {
		if records[i].ID == records[j].ID {
			return records[i].Parent < records[j].Parent
		}
		return records[i].ID < records[j].ID
	})
	var node *Node
	if len(records) == 0 {
		return nil, nil
	}

	if records[0].ID != 0 || records[0].Parent != 0 {
		return nil, fmt.Errorf("bad root record: %v", records[0])
	}

	node = &Node{ID: 0}
	currentID := 0

	for _, record := range records[1:] {

		if record.ID != currentID+1 {
			return nil, fmt.Errorf("non continuous")
		}
		currentID = record.ID

		err := node.Append(&record)
		if err != nil {
			return nil, err
		}
	}

	return node, nil
}
