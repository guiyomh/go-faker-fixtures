package tree

import "github.com/guiyomh/charlatan/pkg/tree/contracts"

// Tree is a binary tree of record
// It sort record with relation
type tree struct {
	Root    contracts.Node
	records map[string]contracts.Node
}

// New create a new binary tree
func New(root contracts.Node) contracts.Tree {
	return &tree{
		Root: root,
		records: map[string]contracts.Node{
			root.Key(): root,
		},
	}
}

// Find and return the node with the supplied key in this subtree. Return nil if not found.
func (me *tree) Find(key string) contracts.Node {
	for k, node := range me.records {
		if k == key {
			return node
		}
	}
	return nil
}

// Add a record in the tree
func (me *tree) Add(node contracts.Node) {
	me.records[node.Key()] = node
	if me.Root != nil {
		me.Root.Add(node)
		return
	}
	me.Root = node
}

// First Return the first (lowest) key and value in the tree, or nil, nil if the tree is empty.
func (me *tree) First() contracts.Node {
	if me.Root == nil {
		return nil
	}
	min := me.Root.Minimum()
	if min == nil {
		return me.Root
	}
	return min
}

// Last Return the last (highest) key and value in the tree, or nil, nil if the tree is empty.
func (me *tree) Last() contracts.Node {
	if me.Root == nil {
		return nil
	}
	max := me.Root.Maximum()
	if max == nil {
		return me.Root
	}
	return max
}

// Walk Iterate the tree with the function in the supplied direction
func (me *tree) Walk(iterator contracts.Iterator, forward bool) {
	if me.Root == nil {
		return
	}
	if forward {
		me.Root.WalkForward(func(node contracts.Node) { iterator(node) })
	} else {
		me.Root.WalkBackward(func(node contracts.Node) { iterator(node) })
	}
}
