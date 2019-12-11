package js_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/dianelooney/kin/pkg/compiler/js"
	"github.com/dianelooney/kin/pkg/parser"
	"github.com/dianelooney/kin/pkg/scanner"
)

const tmpDir = "../../../tmp/"

func prep(src string, t *testing.T) *js.C {
	sc := scanner.NewS(src)
	p := parser.New(sc)
	n, err := p.Parse()
	if err != nil {
		panic(err)
	}
	return js.New(n)
}
func TestDefinitions(t *testing.T) {
	src := `
	def a add "Hello, " "World!"
	log a
	`
	c := prep(src, t)
	out := c.Compile()
	ioutil.WriteFile(tmpDir+t.Name()+".js", []byte(out), os.ModePerm)
	t.Log(out)
}
func TestLogBuiltin(t *testing.T) {
	src := `
	log $ add "Hello, " "World!"`
	c := prep(src, t)
	out := c.Compile()
	ioutil.WriteFile(tmpDir+t.Name()+".js", []byte(out), os.ModePerm)
	t.Log(out)
}

func TestReduceBuiltin(t *testing.T) {
	src := `log $ reduce add 0 [0 1 2 3 4 5 6 7 8 9 10]`
	c := prep(src, t)
	out := c.Compile()
	ioutil.WriteFile(tmpDir+t.Name()+".js", []byte(out), os.ModePerm)
	t.Log(out)
}
func TestFunc(t *testing.T) {
	src := `def something func (a b) add a b`
	c := prep(src, t)
	out := c.Compile()
	ioutil.WriteFile(tmpDir+t.Name()+".js", []byte(out), os.ModePerm)
	t.Log(out)
}
