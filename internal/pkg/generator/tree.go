package generator

import (
	"errors"

	internalcontracts "github.com/guiyomh/charlatan/internal/contracts"
	"github.com/guiyomh/charlatan/pkg/tree"
	treecontracts "github.com/guiyomh/charlatan/pkg/tree/contracts"
)

var (
	// ErrNoRows is returned when the rows collection is empty
	ErrNoRows = errors.New("There are not rows in to build the tree")
)

func BuildTree(rows []internalcontracts.Row) (treecontracts.Tree, error) {
	if len(rows) == 0 {
		return nil, ErrNoRows
	}
	firstRow, _ := rows[0].(treecontracts.Node)
	tree := tree.New(firstRow)
	for _, row := range rows[1:] {
		node, _ := row.(treecontracts.Node)
		tree.Add(node)
	}
	return tree, nil
}
