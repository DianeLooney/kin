package token

type T uint8

const (
	Invalid T = iota
	LParen
	RParen
	LCurly
	RCurly
	LSquare
	RSquare
	Dollar
	Identifier
	Number
	String
	Symbol
)

var lookup = map[byte]T{
	'(': LParen,
	')': RParen,
	'{': LCurly,
	'}': RCurly,
	'[': LSquare,
	']': RSquare,
	'$': Dollar,
}

func Lookup(b byte) (t T, ok bool) {
	t, ok = lookup[b]
	return
}

type Position struct {
	Filename string
	Line     int
	Column   int
}
