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
	N     `json:"-"`
	Name  Node
	Value Node
}
type Function struct {
	N    `json:"-"`
	Args *ArgList
	Body Node
}
type ArgList struct {
	N         `json:"-"`
	Arguments []Node
}

type SExpression struct {
	N        `json:"-"`
	Children []Node
}

type Tag struct {
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
var funcBytes = []byte("func")
var lambdaBytes = []byte("λ")

func (i *Identifier) IsDef() bool {
	return bytes.Equal(i.Raw, defBytes)
}
func (i *Identifier) IsFunc() bool {
	return bytes.Equal(i.Raw, funcBytes)
}
func (i *Identifier) IsLambda() bool {
	return bytes.Equal(i.Raw, lambdaBytes)
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
