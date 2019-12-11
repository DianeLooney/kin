package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/dianelooney/kin/pkg/compiler/js"
	"github.com/dianelooney/kin/pkg/parser"
	"github.com/dianelooney/kin/pkg/scanner"
)

var outFile = flag.String("o", "out.js", "Output file to be read by node")

func main() {
	flag.Parse()
	inFile := flag.Arg(0)
	data, err := ioutil.ReadFile(inFile)
	if err != nil {
		panic(err)
	}
	sc := scanner.New(data)
	p := parser.New(sc)
	n, err := p.Parse()
	if err != nil {
		panic(err)
	}
	out := js.New(n).Compile()
	ioutil.WriteFile(*outFile, []byte(out), os.ModePerm)
}
