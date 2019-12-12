package glsl

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dianelooney/kin/pkg/ast"
)

func nyi(feat string) error {
	return errors.New("NYI - " + feat)
}

func New(n ast.Node) *C {
	return &C{n}
}

type C struct {
	root ast.Node
}

func (c *C) Compile() (out string, err error) {
	return c.render(c.root)
}

func (c *C) render(node ast.Node) (out string, err error) {
	switch n := node.(type) {
	case *ast.Document:
		return c.renderDocument(n)
	case *ast.Number:
		return c.renderNumber(n)
	default:
		return "", nyi(fmt.Sprintf("render(%T)", node))
	}
}

func (c *C) renderDocument(n *ast.Document) (out string, err error) {
	sb := strings.Builder{}
	sb.WriteString("#version 330\n")
	for _, child := range n.Children {
		s, err := c.render(child)
		if err != nil {
			return "", err
		}
		sb.WriteString(s)
		sb.WriteByte('\n')
	}
	out = sb.String()
	return
}

func (c *C) renderNumber(n *ast.Number) (out string, err error) {
	return string(n.Raw), nil
}
