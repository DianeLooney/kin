package parser

import (
	"errors"
	"fmt"
	"io"

	"github.com/dianelooney/kin/pkg/token"

	"github.com/dianelooney/kin/pkg/ast"
	"github.com/dianelooney/kin/pkg/scanner"
)

func New(sc *scanner.Scanner) *Parser {
	return &Parser{
		sc: sc,
	}
}

type Parser struct {
	sc *scanner.Scanner
}

func (p *Parser) Parse() (ast.Node, error) {
	e := ast.Document{}
	// e.Children = append(e.Children, &ast.Reference{Raw: []byte("do")})
	for {
		_, ok := p.sc.Peek()
		if !ok {
			break
		}
		expr, err := p.parseExpression()
		if err != nil {
			return &e, err
		}
		e.Children = append(e.Children, expr)
	}
	return &e, nil
}
func (p *Parser) parseExpression() (ast.Node, error) {
	e := ast.Expression{}

	_, ok := p.sc.Peek()
	if !ok {
		return nil, io.ErrUnexpectedEOF
	}

	first := true
	for {
		if p.sc.HasEOL() && !first {
			break
		}
		b, ok := p.sc.Peek()
		if !ok {
			return nil, io.ErrUnexpectedEOF
		}
		if b == ')' {
			break
		}
		if b == '$' {
			p.sc.Scan()
			e2 := ast.Expression{Children: e.Children}
			e3, err := p.parseExpression()
			if err != nil {
				return nil, err
			}
			e.Children = []ast.Node{&e2, e3}
			return &e, nil
		}
		val, err := p.parseValue()
		if err != nil {
			return nil, err
		}
		if def, ok := val.(*ast.Identifier); first && ok && def.IsDef() {
			d := ast.Definition{}
			ident, err := p.parseValue()
			if err != nil {
				return nil, err
			}
			d.Identifier, ok = ident.(*ast.Identifier)
			if !ok {
				return nil, errors.New("Expected identifier")
			}
			e, err := p.parseExpression()
			if err != nil {
				return nil, err
			}
			d.Value = e
			return &d, nil
		}
		if fun, ok := val.(*ast.Identifier); first && ok && fun.IsFunc() {
			f := ast.Function{}
			args, err := p.parseArgumentList()
			if err != nil {
				return nil, err
			}
			f.Args = args.(*ast.ArgList)
			body, err := p.parseExpression()
			if err != nil {
				return nil, err
			}
			f.Body = body
			return &f, nil
		}
		e.Children = append(e.Children, val)
		first = false
	}
	return &e, nil
}
func (p *Parser) parseValue() (ast.Node, error) {
	b, ok := p.sc.Peek()
	if !ok {
		return nil, io.ErrUnexpectedEOF
	}
	switch b {
	case '(':
		return p.parseSExpression()
	case '{':
		return p.parseObject()
	case '[':
		return p.parseArray()
	}

	lit, _, t, err := p.sc.Scan()
	if err != nil {
		return nil, err
	}
	if t == token.String {
		return &ast.String{Raw: lit}, nil
	}
	if t == token.Symbol {
		return &ast.Symbol{Raw: lit}, nil
	}
	if t == token.Number {
		return &ast.Number{Raw: lit}, nil
	}
	if t == token.Identifier {
		return &ast.Identifier{Raw: lit}, nil
	}

	return nil, fmt.Errorf("NYI - %s", lit)
}
func (p *Parser) parseArgumentList() (ast.Node, error) {
	_, _, t, err := p.sc.Scan()
	if err != nil {
		return nil, err
	}
	if t != token.LParen {
		return nil, errors.New("Expected '(' to start argument list")
	}

	n := ast.ArgList{}
	for {
		b, ok := p.sc.Peek()
		if !ok {
			return nil, io.ErrUnexpectedEOF
		}
		if b == ')' {
			p.sc.Scan()
			break
		}
		v, err := p.parseValue()
		if err != nil {
			return nil, err
		}
		n.Arguments = append(n.Arguments, v)
	}
	return &n, nil
}
func (p *Parser) parseSExpression() (ast.Node, error) {
	_, _, t, err := p.sc.Scan()
	if err != nil {
		return nil, err
	}
	if t != token.LParen {
		return nil, errors.New("Expected '(' to start array definition")
	}

	n := ast.SExpression{}
	for {
		b, ok := p.sc.Peek()
		if !ok {
			return nil, io.ErrUnexpectedEOF
		}
		if b == ')' {
			p.sc.Scan()
			break
		}
		v, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		n.Children = append(n.Children, v)
	}
	return &n, nil
}
func (p *Parser) parseArray() (ast.Node, error) {
	_, _, t, err := p.sc.Scan()
	if err != nil {
		return nil, err
	}
	if t != token.LSquare {
		return nil, errors.New("Expected '[' to start array definition")
	}

	n := ast.Array{}
	for {
		b, ok := p.sc.Peek()
		if !ok {
			return nil, io.ErrUnexpectedEOF
		}
		if b == ']' {
			p.sc.Scan()
			break
		}
		v, err := p.parseValue()
		if err != nil {
			return nil, err
		}
		n.Values = append(n.Values, v)
	}
	return &n, nil
}
func (p *Parser) parseObject() (ast.Node, error) {
	_, _, t, err := p.sc.Scan()
	if err != nil {
		return nil, err
	}
	if t != token.LCurly {
		return nil, errors.New("Expected '{' to start object definition")
	}

	n := ast.Object{}
	for {
		b, ok := p.sc.Peek()
		if !ok {
			return nil, io.ErrUnexpectedEOF
		}
		if b == '}' {
			p.sc.Scan()
			break
		}
		k, err := p.parseValue()
		if err != nil {
			return nil, err
		}
		v, err := p.parseValue()
		if err != nil {
			return nil, err
		}
		n.Pairs = append(n.Pairs, &ast.ObjectPair{
			Key:   k,
			Value: v,
		})
	}
	return &n, nil
}
