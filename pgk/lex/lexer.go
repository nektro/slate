package lex

import (
	"log"
)

type Lexer struct {
	Inner
	l int
	p int
}

type Inner interface {
	CouldBeWord(string) bool
	IsKeyword(string) bool
	IsNumber(string) bool
	IsSymbol(rune) bool
	IsQuote(rune) bool
	SLineComment() rune
}

var (
	skippedChars = []rune{
		' ', '\n',
	}
)

func isSkippedRune(r rune) bool {
	switch r {
	case ' ', '\n', '\t', '\r':
		return true
	}
	return false
}

func (v *Lexer) Lex(in string, f func(*Token)) {
	v.l = 1
	v.p = 1

	s := 0
	e := 0
	m := 0
	for i := 0; i < len(in); i++ {
		c := rune(in[i])

		flsh := func() bool {
			if m == 0 {
				if c == v.SLineComment() {
					m = 1
					return false
				}
				if v.IsQuote(c) {
					m = 2
					return false
				}
			}
			if m == 1 {
				ecm := c == '\n'
				if ecm {
					f(v.handle(TTCom, in[s:i]))
					s = i
					e = i
					m = 0
				}
				return ecm
			}
			if m == 2 {
				if c == rune(in[s]) {
					f(v.handle(TTStr, in[s:i+1]))
					s = i + 1
					e = i
					m = 0
				}
				return false
			}
			if isSkippedRune(c) {
				return true
			}
			if v.IsSymbol(c) {
				return true
			}
			return false
		}()

		func() {
			if !flsh {
				e++
			}
			if flsh {
				if m == 0 {
					if e-s > 0 {
						f(v.handle(TTWrd, in[s:e]))
						s = i
						e = i
					}
					if isSkippedRune(c) {
						s++
						e++
					}
					if v.IsSymbol(c) {
						f(v.handle(TTSym, string(c)))
						s++
						e++
					}
				}
			}
		}()

		v.p++
		if c != '\n' {
			continue
		}
		v.l++
		v.p = 1
	}
}

func (v *Lexer) handle(t TokenType, s string) (r *Token) {
	r = &Token{t, s, v.l, v.p - len(s), len(s)}

	if r.T != TTWrd {
		return
	}
	if v.IsKeyword(r.V) {
		r.T = TTKey
		return
	}
	if v.IsNumber(r.V) {
		r.T = TTNum
		return
	}
	if v.CouldBeWord(r.V) {
		return
	}
	log.Println("lex: bad word:", r.V, r.LP())
	log.Println(r)
	log.Fatalln()
	return
}

func (v *Lexer) LexAll(filename, in string) *File {
	res := []*Token{}
	v.Lex(in, func(t *Token) {
		res = append(res, t)
	})
	return &File{file: filename, tokens: res}
}
