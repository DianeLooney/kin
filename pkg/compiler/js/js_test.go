package js_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/dianelooney/kin/pkg/compiler/js"
	"github.com/dianelooney/kin/pkg/parser"
	"github.com/dianelooney/kin/pkg/scanner"
)

func prep(src string) *js.C {
	sc := scanner.NewS(src)
	p := parser.New(sc)
	n, err := p.Parse()
	if err != nil {
		panic(err)
	}
	return js.New(n)
}

func TestLogBuiltin(t *testing.T) {
	src := `
	log $ add "Hello, " "World!"`
	c := prep(src)
	out := c.Compile()
	ioutil.WriteFile(t.Name(), []byte(out), os.ModePerm)
	t.Log(out)
}
