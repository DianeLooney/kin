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
	return fmt.Sprintf(runtime, c.render(c.root))
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
	case *ast.Number:
		return c.renderNumber(n)
	case *ast.Array:
		return c.renderArray(n)
	case *ast.SExpression:
		return c.renderSExpression(n)
	case *ast.Function:
		return c.renderFunction(n)
	default:
		return fmt.Sprintf("/* NYI - %T */", n)
	}
}
func (c *C) renderDocument(n *ast.Document) (out string) {
	var strs = make([]string, len(n.Children))
	for i, child := range n.Children {
		strs[i] = c.render(child)
	}
	return strings.Join(strs, ";\n")
}
func (c *C) renderDefinition(n *ast.Definition) (out string) {
	const tmpl = `const %v = %v;`
	return fmt.Sprintf(tmpl, c.render(n.Name), c.render(n.Value))
}
func (c *C) renderExpression(n *ast.Expression) (out string) {
	if len(n.Children) == 1 {
		return c.render(n.Children[0])
	}

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
func (c *C) renderArray(n *ast.Array) (out string) {
	strs := make([]string, len(n.Values))
	for i, v := range n.Values {
		strs[i] = c.render(v)
	}
	return fmt.Sprintf("[%v]", strings.Join(strs, ", "))
}
func (c *C) renderNumber(n *ast.Number) (out string) {
	return string(n.Raw)
}
func (c *C) renderFunction(n *ast.Function) (out string) {
	const tmpl = `(%v => %v)`
	args := make([]string, len(n.Args.Arguments))
	for i, arg := range n.Args.Arguments {
		args[i] = c.render(arg)
	}
	body := c.render(n.Body)
	return fmt.Sprintf(tmpl, strings.Join(args, " => "), body)
}
func (c *C) renderSExpression(n *ast.SExpression) (out string) {
	const template = `(()=>{return %v})()`
	var strs []string
	for _, child := range n.Children {
		strs = append(strs, c.render(child))
	}
	return strings.Join(strs, ", ")
}
