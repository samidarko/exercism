package pov

type Tree struct {
	// Add the needed fields here
	value    string
	children []*Tree
	parent   *Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	tree := &Tree{value: value, children: children, parent: nil}
	for _, child := range children {
		child.parent = tree
	}
	return tree
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

func (tr *Tree) Siblings(sibling *Tree) []*Tree {
	siblings := make([]*Tree, 0)
	for _, child := range tr.children {
		if child.Value() != sibling.Value() {
			siblings = append(siblings, child)
		}
	}
	return siblings
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {

	node := tr.FindNode(from)
	if node == nil {
		return nil
	}
	t, p := node, node.parent

	node.parent = nil

	for p != nil {
		t.children = append(t.children, p)
		k := 0

		for i, c := range p.children {
			if c == t {
				k = i
				break
			}
		}

		p.children = append(p.children[:k], p.children[k+1:]...)
		q := p.parent
		p.parent = t
		t, p = p, q
	}

	return node
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {

	node := tr.FromPov(from)
	if node == nil {
		return nil
	}

	type Element struct {
		node *Tree
		path []string
	}

	queue := []Element{{node, []string{}}}

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]
		element.path = append(element.path, element.node.Value())
		if element.node.Value() == to {
			return element.path
		}
		for _, child := range element.node.Children() {
			queue = append(queue, Element{child, element.path})
		}
	}

	return nil
}

// FindNode returns the node or nil
func (tr *Tree) FindNode(value string) *Tree {
	if tr.Value() == value {
		return tr
	}
	for _, child := range tr.Children() {
		if node := child.FindNode(value); node != nil {
			return node
		}
	}
	return nil
}
