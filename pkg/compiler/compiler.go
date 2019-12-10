package compiler

import (
	"github.com/dianelooney/kin/pkg/ast"
)

type Compiler interface {
	Compile(n ast.Node) (output []byte, err error)
}
