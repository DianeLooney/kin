package ast

import (
	"bytes"

	"github.com/dianelooney/kin/pkg/token"
)

type Node interface {
	Position() token.Position
}

type N struct {
	P token.Position
}

func (n *N) Position() token.Position {
	return n.P
}

type Document struct {
	N        `json:"-"`
	Children []Node
}

type Expression struct {
	N        `json:"-"`
	Children []Node
}

type Definition struct {
	N          `json:"-"`
	Identifier *Identifier
	Value      Node
}

type SExpression struct {
	N        `json:"-"`
	Children []Node
}

type Reference struct {
	N   `json:"-"`
	Raw []byte
}

type Number struct {
	N   `json:"-"`
	Raw []byte `json:","`
}

type Identifier struct {
	N   `json:"-"`
	Raw []byte
}

var defBytes = []byte("def")

func (i *Identifier) IsDef() bool {
	return bytes.Equal(i.Raw, defBytes)
}

type Symbol struct {
	N   `json:"-"`
	Raw []byte
}

type String struct {
	N   `json:"-"`
	Raw []byte
}

type Array struct {
	N      `json:"-"`
	Values []Node
}

type Object struct {
	N     `json:"-"`
	Pairs []Node
}

type ObjectPair struct {
	N     `json:"-"`
	Key   Node
	Value Node
}
