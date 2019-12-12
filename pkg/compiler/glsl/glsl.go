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
	case *ast.Definition:
		return c.renderDefinition(n)
	case *ast.Document:
		return c.renderDocument(n)
	case *ast.Function:
		return c.renderFunction(n)
	case *ast.Identifier:
		return c.renderIdentifier(n)
	case *ast.Number:
		return c.renderNumber(n)
	default:
		return "", nyi(fmt.Sprintf("render(%T)", node))
	}
}

func (c *C) renderDefinition(n *ast.Definition) (out string, err error) {
	tn, ok := n.Name.(*ast.Tag)
	if !ok {
		return "", fmt.Errorf("Expected definition's name to be a typed s-expression, but it was a '%T'", n.Name)
	}
	if len(tn.Children) != 2 {
		return "", fmt.Errorf("Expected definitions name to be length 2, (name type), but it had %v values", len(tn.Children))
	}
	name, ok := tn.Children[0].(*ast.Identifier)
	if !ok {
		return "", fmt.Errorf("Expected definitions name to be two identifiers, but the first was a %T", tn.Children[0])
	}
	typ, ok := tn.Children[1].(*ast.Identifier)
	if !ok {
		return "", fmt.Errorf("Expected definitions name to be two identifiers, but the second was a %T", tn.Children[0])
	}
	nameR, err := c.render(name)
	if err != nil {
		return "", err
	}
	typeR, err := c.render(typ)
	if err != nil {
		return "", err
	}
	body, err := c.render(n.Value)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s %s;", typeR, nameR, body), nil
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

func (c *C) renderIdentifier(n *ast.Identifier) (out string, err error) {
	return string(n.Raw), nil
}

const fnTemplate = `(%s) {
%v
}`

func (c *C) renderFunction(n *ast.Function) (out string, err error) {
	args := make([]string, len(n.Args.Arguments))
	for i, arg := range n.Args.Arguments {
		args[i], err = c.render(arg)
		if err != nil {
			return
		}
	}
	switch n.Body {
	case *ast.Expression:

	case *ast.SExpression:
	}
	body, err := c.render(n.Body)
	if err != nil {
		return
	}
	out = fmt.Sprintf(fnTemplate, strings.Join(args, ", "), body)
	return
}

func (c *C) renderNumber(n *ast.Number) (out string, err error) {
	return string(n.Raw), nil
}

func (c *C) renderTag(n *ast.Tag) (out string, err error) {
	if len(n.Children) != 2 {
		return "", fmt.Errorf("Expected tag to have two values, but it had %v", len(n.Children))
	}
	name, err := c.render(n.Children[0])
	if err != nil {
		return
	}
	typ, err := c.render(n.Children[1])
	if err != nil {
		return
	}
	out = fmt.Sprintf("%v %v", typ, name)
	return
}
