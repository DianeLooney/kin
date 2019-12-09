package scanner_test

import (
	"io"
	"strings"
	"testing"

	"github.com/dianelooney/kin/pkg/token"

	"github.com/dianelooney/kin/pkg/scanner"
)

func TestSingleToken(t *testing.T) {
	s := scanner.NewS(`()`)
	s.Scan()
}
func TestSkipWhitespace(t *testing.T) {
	cases := map[string]string{
		"spaces": ` ()`,
		"tabs": `	()`,
		"newline": `
		()`,
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			s := scanner.NewS(c)
			{
				_, _, tok, err := s.Scan()
				if err != nil {
					t.Errorf("Expected no error, but got %v", err)
				}
				if tok != token.LParen {
					t.Errorf("Expected LParen (%v) but got %v", token.LParen, tok)
				}
			}
			{
				_, _, tok, err := s.Scan()
				if err != nil {
					t.Errorf("Expected no error, but got %v", err)
				}
				if tok != token.RParen {
					t.Errorf("Expected RParen (%v) but got %v", token.RParen, tok)
				}
			}
		})
	}
}

func TestNumbers(t *testing.T) {
	cases := []string{
		"1",
		"-1",
		"10.",
		"-10.",
		"10.23",
		"-10.23",
		"0",
		"-0",
		".4",
		"-.4",
		" 420",
		"-69 ",
	}
	for _, c := range cases {
		t.Run(c, func(t *testing.T) {
			s := scanner.NewS(c)
			lit, _, tok, err := s.Scan()
			if err != nil {
				t.Errorf("Expected no error, but received %v", err)
			}
			if string(lit) != strings.Trim(c, " \t") {
				t.Errorf("Expected scanned token to be '%s', but it was '%s'", c, lit)
			}
			if tok != token.Number {
				t.Errorf("Expected to receive a Number (%v), but it was %v", token.Number, tok)
			}
		})
	}
}

func TestStrings(t *testing.T) {
	cases := []string{
		`""`,
		`"\""`,
		`"\'" `,
		` "
		"`,
	}
	for _, c := range cases {
		t.Run(c, func(t *testing.T) {
			s := scanner.NewS(c)
			lit, _, tok, err := s.Scan()
			if err != nil {
				t.Errorf("Expected no error, but received %v", err)
			}
			if string(lit) != strings.Trim(c, " \t") {
				t.Errorf("Expected scanned token to be '%s', but it was '%s'", c, lit)
			}
			if tok != token.String {
				t.Errorf("Expected to receive a Number (%v), but it was %v", token.Number, tok)
			}
		})
	}
}
func TestSymbols(t *testing.T) {
	cases := []string{
		`:asdf`,
		`:asdf'something`,
		`:asdf"something`,
		`:asdf+=something`,
		`:--`,
		`:::`,
		` :asdf`,
		`:asdf `,
		`:🍆`,
		`:يونيكود`,
	}
	for _, c := range cases {
		t.Run(c, func(t *testing.T) {
			s := scanner.NewS(c)
			lit, _, tok, err := s.Scan()
			if err != nil {
				t.Errorf("Expected no error, but received %v", err)
			}
			if string(lit) != strings.Trim(c, " \t") {
				t.Errorf("Expected scanned token to be '%s', but it was '%s'", c, lit)
			}
			if tok != token.Symbol {
				t.Errorf("Expected to receive a Number (%v), but it was %v", token.Number, tok)
			}
		})
	}
}
func TestIdentifiers(t *testing.T) {
	cases := []string{
		`asdf`,
		`asdf'something`,
		`asdf"something`,
		`asdf+=something`,
		`--`,
		` asdf`,
		`asdf `,
		`🍆`,
		`يونيكود`,
	}
	for _, c := range cases {
		t.Run(c, func(t *testing.T) {
			s := scanner.NewS(c)
			lit, _, tok, err := s.Scan()
			if err != nil {
				t.Errorf("Expected no error, but received %v", err)
			}
			if string(lit) != strings.Trim(c, " \t") {
				t.Errorf("Expected scanned token to be '%s', but it was '%s'", c, lit)
			}
			if tok != token.Identifier {
				t.Errorf("Expected to receive a Number (%v), but it was %v", token.Number, tok)
			}
		})
	}
}
func TestHasEOL(t *testing.T) {
	src := `x
	asdf
	asdf wxyz
	( do	something )
	
	cool`
	samples := []bool{
		false, //           -> x
		true,  // x         -> \n
		true,  // asdf      -> \n
		false, // asdf      -> wxyz
		true,  // wxyz      -> \n
		false, // (         -> do
		false, // do        -> something
		false, // something -> )
		true,  // )         -> \n
		true,  // cool      -> EOF
	}
	sc := scanner.NewS(src)
	for i, s := range samples {
		ok := sc.HasEOL()
		if ok != s {
			t.Errorf("Sample %v should be %v but it was %v", i, s, ok)
		}
		lit, _, _, err := sc.Scan()
		if err != nil && err != io.EOF {
			t.Errorf("Unable to scan: %v", err)
		}
		t.Log(string(lit))
	}
}
func TestEOF(t *testing.T) {
	s := scanner.NewS(``)
	_, _, _, err := s.Scan()
	if err != io.EOF {
		t.Fail()
	}
}
