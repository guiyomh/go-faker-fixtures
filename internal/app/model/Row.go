package model

import (
	internal "github.com/guiyomh/charlatan/internal/contracts"
	tree "github.com/guiyomh/charlatan/pkg/tree/contracts"
)

// Row represente data that will be persist in database
type row struct {
	tree.Node
	name                string
	tableName           string
	fields              map[string]interface{}
	dependencyReference map[string]internal.Relation
	left                tree.Node
	right               tree.Node
	pk                  interface{}
}

// NewRow Factory to build a Row Structure
func NewRow(name, tableName string) internal.Row {
	return &row{
		name:                name,
		tableName:           tableName,
		fields:              make(map[string]interface{}),
		dependencyReference: make(map[string]internal.Relation, 0),
		pk:                  nil,
	}
}

func (r row) Schema() string {
	return "fixtures" //TODO implement schema
}

func (r row) Pk() interface{} {
	return r.pk
}

func (r row) SetPk(value interface{}) {
	r.pk = value
}

func (r row) Name() string {
	return r.name
}

func (r row) TableName() string {
	return r.tableName
}

func (r row) Fields() map[string]interface{} {
	return r.fields
}

func (r row) AddField(key string, value interface{}) {
	r.fields[key] = value
}

func (r row) AddDependency(key string, value internal.Relation) {
	r.dependencyReference[key] = value
}

func (r row) DependencyReference() map[string]internal.Relation {
	return r.dependencyReference
}

func (r row) hasDependencyOf(name string) bool {
	if !r.HasDependencies() {
		return false
	}
	for _, v := range r.dependencyReference {
		if v.RecordName() == name {
			return true
		}
	}
	return false
}

// HasDependencies return true if this row has dependancies
func (r row) HasDependencies() bool {
	return len(r.dependencyReference) > 0
}

// SetDependance Set the value of a depandance
// func (r row) SetDependance(name string, row *Row) {
// 	r.Fields[name] = value
// }

// LessThan return true if this row hasn't dependancy with other.
func (r row) LessThan(other tree.Node) bool {
	original, ok := other.(row)
	if !ok {
		return false
	}
	return !r.hasDependencyOf(original.Name())
}

// EqualTo return true if this row is equal to the other.
func (r row) EqualTo(other tree.Node) bool {
	original, ok := other.(row)
	if !ok {
		return false
	}
	return r.Name() == original.Name()
}

// GreaterThan return true if this row has dependancy of other.
func (r row) GreaterThan(other tree.Node) bool {
	original, ok := other.(row)
	if !ok {
		return false
	}
	return r.hasDependencyOf(original.Name())
}

// Add an existing node to this node's subtree
func (r row) Add(node tree.Node) {
	if node.LessThan(r) {
		if r.left == nil {
			r.left = node
		}
		r.left.Add(node)
	} else {
		if r.right == nil {
			r.right = node
		}
		r.right.Add(node)
	}
}

// Minimum Return the left-most (smallest key) node in this node's subtree
func (r row) Minimum() tree.Node {
	for {
		if r.left == nil {
			return r
		}
		r = r.left.(row)
	}
}

// Maximum Return the right-most (largest key) node in this node's subtree
func (r row) Maximum() tree.Node {
	for {
		if r.right == nil {
			return r
		}
		r = r.right.(row)
	}
}

// WalkForward Call iterator for each node in this node's subtree in order, low to high
func (r row) WalkForward(iterator tree.Iterator) {
	if r.left != nil {
		r.left.WalkForward(iterator)
	}
	iterator(r)
	if r.right != nil {
		r.right.WalkForward(iterator)
	}
}

// WalkBackward Call iterator for each node in this node's subtree in reverse order, high to low
func (r row) WalkBackward(iterator tree.Iterator) {
	if r.right != nil {
		r.right.WalkBackward(iterator)
	}
	iterator(r)
	if r.left != nil {
		r.left.WalkBackward(iterator)
	}
}

func (r row) Key() string {
	return r.Name()
}
