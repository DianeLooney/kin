package js_test

import (
	"encoding/json"
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
	out, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(tmpDir+t.Name()+".ast.json", out, os.ModePerm)
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
	ioutil.WriteFile(t.Name(), []byte(out), os.ModePerm)
	t.Log(out)
}
