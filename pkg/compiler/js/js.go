package js

import (
	"fmt"
	"strings"

	"github.com/dianelooney/kin/pkg/ast"
)

func New(n ast.Node) *C {
	return &C{n}
}

type C struct {
	root ast.Node
}

func (c *C) Compile() (out string) {
	return runtime + c.render(c.root)
}
func (c *C) render(node ast.Node) (out string) {
	switch n := node.(type) {
	case *ast.Expression:
		return c.renderExpression(n)
	case *ast.Identifier:
		return c.renderIdentifier(n)
	case *ast.String:
		return c.renderString(n)
	case *ast.SExpression:
		return c.renderSExpression(n)
	default:
		return fmt.Sprintf("/* NYI - %T */", n)
	}

}
func (c *C) renderExpression(n *ast.Expression) (out string) {
	for _, child := range n.Children {
		out = out + "(" + c.render(child) + ")"
	}
	return out
}
func (c *C) renderIdentifier(n *ast.Identifier) (out string) {
	return string(n.Raw)
}
func (c *C) renderString(n *ast.String) (out string) {
	return string(n.Raw)
}
func (c *C) renderSExpression(n *ast.SExpression) (out string) {
	var strs []string
	for _, child := range n.Children {
		strs = append(strs, c.render(child))
	}
	return strings.Join(strs, ",")
}

var runtime = `
// BEGIN RUNTIME
const log = console.log;
const add = x => y => x + y;
// END RUNTIME

`
