package contracts

type Node interface {
	LessThan(other Node) bool
	EqualTo(other Node) bool
	GreaterThan(other Node) bool
	Add(node Node)
	Minimum() Node
	Maximum() Node
	WalkForward(iterator Iterator)
	WalkBackward(iterator Iterator)
	Key() string
}
