package tree

import (
	"testing"

	mocks "github.com/guiyomh/charlatan/pkg/tree/contracts/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	root := &mocks.Node{}
	root.On("Minimum").Return(root)
	root.On("Key").Return("bar")
	tree := New(root)
	assert.Equal(t, root, tree.First())
}

func Test_tree_Add(t *testing.T) {

	// Create a node
	node := &mocks.Node{}
	node.On("Key").Return("foo_node")

	//Create a root node
	root := &mocks.Node{}
	root.On("Add", node)
	root.On("Key").Return("foo_root")

	//Create tree
	tree := New(root)
	tree.Add(node)
}

func Test_tree_First(t *testing.T) {
	//Test tree without Node and Root
	tree := tree{}
	assert.Nil(t, tree.First())

	//Test tree only with Root
	root := &mocks.Node{}
	root.On("Minimum").Return(nil)
	root.On("Key").Return("bar_root")

	treeWithRoot := New(root)
	assert.Equal(t, root, treeWithRoot.First())

	// Test tree with root and Node
	root2 := &mocks.Node{}
	root2.On("Key").Return("foo_root")
	node := &mocks.Node{}
	root2.On("Minimum").Return(node)

	treeWithNode := New(root2)

	assert.Equal(t, node, treeWithNode.First())
}

func Test_tree_Last(t *testing.T) {
	//Test tree without Node and Root
	tree := tree{}
	assert.Nil(t, tree.Last())

	//Test tree only with Root
	root := &mocks.Node{}
	root.On("Maximum").Return(nil)
	root.On("Key").Return("foo_root")

	treeWithRoot := New(root)
	assert.Equal(t, root, treeWithRoot.Last())

	// Test tree with root and Node
	root2 := &mocks.Node{}
	node := &mocks.Node{}
	root2.On("Maximum").Return(node)
	root2.On("Key").Return("bar_root")

	treeWithNode := New(root2)

	assert.Equal(t, node, treeWithNode.Last())
}
