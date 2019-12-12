package glsl_test

import (
	"strings"
	"testing"

	. "github.com/go-playground/assert/v2"

	"github.com/dianelooney/kin/pkg/compiler/glsl"
	"github.com/dianelooney/kin/pkg/parser"
	"github.com/dianelooney/kin/pkg/scanner"
)

func compileS(in string, t *testing.T) (out string) {
	sc := scanner.NewS(in)
	p := parser.New(sc)
	n, err := p.Parse()
	if err != nil {
		t.Errorf("Unable to parse example code: %v", err)
		return
	}
	out, err = glsl.New(n).Compile()
	if err != nil {
		t.Errorf("Unable to compile example code: %v", err)
	}
	return
}

func EqualTrimmed(t *testing.T, s1 string, s2 string) {
	s1 = strings.Trim(s1, " \t\n")
	s2 = strings.Trim(s2, " \t\n")
	Equal(t, s1, s2)
}
func testIt(t *testing.T, in string, expected string) {
	out := compileS(in, t)
	EqualTrimmed(t, out, expected)
}
func TestCanRender(t *testing.T) {
	in := ``
	out := compileS(in, t)
	expected := `#version 330`
	EqualTrimmed(t, out, expected)
}

const mainSrc = `
def <main void> func () set outputColor 0
`
const mainFix = `
#version330
void main() {
};
`

func TestCanRenderMain(t *testing.T) {
	testIt(t, mainSrc, mainFix)
}
