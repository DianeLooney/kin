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
	case *ast.Document:
		return c.renderDocument(n)
	case *ast.Definition:
		return c.renderDefinition(n)
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
func (c *C) renderDocument(n *ast.Document) (out string) {
	var strs = make([]string, len(n.Children))
	for i, child := range n.Children {
		strs[i] = c.render(child)
	}
	return strings.Join(strs, "\n")
}
func (c *C) renderDefinition(n *ast.Definition) (out string) {
	const tmpl = `const %v = %v;`
	return fmt.Sprintf(tmpl, c.render(n.Identifier), c.render(n.Value))
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
	const template = `(()=>{return %v})()`
	var strs []string
	for _, child := range n.Children {
		strs = append(strs, c.render(child))
	}
	return strings.Join(strs, ", ")
}

var runtime = `

// BEGIN RUNTIME
const log = console.log;
const add = x => y => x + y;
const head = x => x[0];
const tail = x => x.slice(1);
const init = x => x.slice(0, x.length - 1);
const last = x => x[x.length];
// END RUNTIME

`
