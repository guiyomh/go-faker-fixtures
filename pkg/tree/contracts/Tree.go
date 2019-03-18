package contracts

// Tree represents a public api for a binary tree
type Tree interface {
	Add(node Node)
	First() Node
	Last() Node
	Find(string) Node
	Walk(iterator Iterator, forward bool)
}
