package scanner

import (
	"fmt"
	"io"
	"regexp"

	"github.com/dianelooney/kin/pkg/token"
)

func New(source []byte) (sc *Scanner) {
	return &Scanner{
		source: source,
		line:   0,
		char:   0,
	}
}

func NewS(source string) (sc *Scanner) {
	return New([]byte(source))
}

type Scanner struct {
	source []byte
	line   int
	char   int
}

func (s *Scanner) HasEOL() bool {
	for _, c := range s.source {
		switch c {
		case '\t':
			continue
		case ' ':
			continue
		case '\n':
			return true
		default:
			fmt.Printf("%s\n", []byte{c})
			return false
		}
	}
	return true
}

func (s *Scanner) Scan() (literal []byte, p token.Position, t token.T, err error) {
	s.skipWhitespace()

	if len(s.source) == 0 {
		return nil, token.Position{}, token.Invalid, io.EOF
	}

	if tok, ok := token.Lookup(s.source[0]); ok {
		literal, s.source = s.source[0:1], s.source[1:]
		p = token.Position{
			Line:   s.line,
			Column: s.char,
		}
		t = tok
		return
	}

	if s.canScan(numRegexp) {
		return s.scan(numRegexp, token.Number)
	}
	if s.canScan(strRegexp) {
		return s.scan(strRegexp, token.String)
	}
	if s.canScan(symRegexp) {
		return s.scan(symRegexp, token.Symbol)
	}
	return s.scan(identRegexp, token.Identifier)
}

func (s *Scanner) skipWhitespace() {
	for len(s.source) > 0 {
		switch s.source[0] {
		case ' ':
			fallthrough
		case '\t':
			s.char++
		case '\n':
			s.line++
			s.char = 0
		case '\r':
		default:
			return
		}
		s.source = s.source[1:]
	}
}

func (s *Scanner) canScan(r *regexp.Regexp) (ok bool) {
	return r.Match(s.source)
}

func (s *Scanner) scan(r *regexp.Regexp, tok token.T) (literal []byte, p token.Position, t token.T, err error) {
	match := r.FindSubmatch(s.source)
	literal = match[0]
	p = token.Position{
		Line:   s.line,
		Column: s.char,
	}
	t = tok
	err = nil

	for _, c := range literal {
		if c == '\n' {
			s.line++
			s.char = 0
		} else {
			s.char++
		}
	}
	s.source = s.source[len(match[0]):]
	return
}

var numRegexp = regexp.MustCompile(`^[+-]?((\d+(\.\d*)?)|(\.\d+))`)
var strRegexp = regexp.MustCompile(`^"(?:[^"\\]|\\.)*"`)
var symRegexp = regexp.MustCompile(`^:[^\s\()[\]{}]+`)
var identRegexp = regexp.MustCompile(`^[^\s\()[\]{}]+`)
