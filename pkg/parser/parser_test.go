package parser_test

import (
	"testing"

	"github.com/dianelooney/kin/pkg/parser"
	"github.com/dianelooney/kin/pkg/scanner"
)

func TestParse(t *testing.T) {
	srcs := map[string]string{
		"basic":         `do something`,
		"symbols":       `do :something`,
		"strings":       `do "something" cool`,
		"numbers":       `do something 42.0 69`,
		"arrays":        `map asdf [420 69]`,
		"empty objects": `{}`,
		"objects":       `merge {:a 1 :b "bananas"} {:c 2 :b "bananarama"}`,
		"s-exprs":       `do (something cool)`,
		"dollar":        `a 4.0 $ b 3.0`,
	}
	for n, src := range srcs {
		t.Run(n, func(t *testing.T) {
			sc := scanner.NewS(src)
			p := parser.New(sc)
			_, err := p.Parse()
			if err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}
		})
	}
}
