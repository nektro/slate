package slate

import (
	"strings"

	"github.com/nektro/go-util/arrays/stringsu"
	"github.com/nektro/slate/pgk/lex"
)

type sLexer struct {
	k []string
	s string
	q string
}

// singleton
var (
	Lexer = &lex.Lexer{
		Inner: &sLexer{
			[]string{
				"pub", "const", "asm", "true", "false",
				"func", "return", "defer", "errdefer", "comptime",
				"if", "else", "for", "break", "continue",
				"union", "enum", "struct", "interface",
			},
			`(){}[],.;=+-*/%<>|&^!@:`,
			`"'`,
		},
	}
)

func (v *sLexer) CouldBeWord(s string) bool {
	b := 0
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			continue
		}
		if c >= 'A' && c <= 'Z' {
			continue
		}
		if c >= '0' && c <= '9' {
			continue
		}
		if c == '_' {
			continue
		}
		b++
	}
	if b > 0 {
		return false
	}
	if s[0] >= '0' && s[0] <= '9' {
		return false
	}
	return true
}

func (v *sLexer) IsKeyword(s string) bool {
	return stringsu.Contains(v.k, s)
}

func (v *sLexer) IsNumber(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func (v *sLexer) IsSymbol(r rune) bool {
	return strings.ContainsRune(v.s, r)
}

func (v *sLexer) IsQuote(r rune) bool {
	return strings.ContainsRune(v.q, r)
}

func (v *sLexer) SLineComment() rune {
	return '#'
}
