package ast

import (
	"github.com/dianelooney/kin/pkg/token"
)

type Node interface {
	Position() token.Position
}

type node struct {
	P token.Position
}

func (n *node) Position() token.Position {
	return n.P
}

type Expression struct {
	node
	Children []Node
}
type SExpression struct {
	node
	Children []Node
}

type Reference struct {
	node
	Raw []byte
}

type Number struct {
	node
	Raw []byte
}

type Identifier struct {
	node
	Raw []byte
}

type Symbol struct {
	node
	Raw []byte
}

type String struct {
	node
	Raw []byte
}

type Array struct {
	node
	Values []Node
}

type Object struct {
	node
	Pairs []Node
}

type ObjectPair struct {
	node
	Key   Node
	Value Node
}
