package lex

import (
	"fmt"
	"log"
)

type File struct {
	file   string
	tokens []*Token
	index  int
}

func (f *File) Len() int {
	return len(f.tokens)
}

func (f *File) RemoveComments() {
	res := []*Token{}
	for _, item := range f.tokens {
		if item.T != TTCom {
			res = append(res, item)
		}
	}
	f.tokens = res
}

func (f *File) Print() {
	for i, item := range f.tokens {
		log.Println(f.file+":"+item.LP(), i, item.PL())
	}
}

func (f *File) get() *Token {
	if f.index >= len(f.tokens) {
		panic(fmt.Sprint("parse failure:", "specified token index at or past file length"))
	}
	t := f.tokens[f.index]
	return t
}

func (f *File) ret(t *Token) *Token {
	f.index++
	return t
}

func (f *File) fail(s, v string, t *Token) {
	panic(fmt.Sprint("expected token:", s+":", v, ",", "instead got:", t.PL(), ":", f.file+":"+t.LP()))
}

func (f *File) Left() []*Token {
	return f.tokens[f.index:]
}

func (f *File) Type(y TokenType) *Token {
	t := f.get()
	if t.T != y {
		panic(fmt.Sprint("expected token type:", y, ",", "instead got:", t.PL(), ":", f.file+":"+t.LP()))
	}
	return f.ret(t)
}

func (f *File) TypeS(y string) *Token {
	t := f.get()
	if string(t.T) != y {
		panic(fmt.Sprint("expected token type:", y, ",", "instead got:", t.PL(), ":", f.file+":"+t.LP()))
	}
	return f.ret(t)
}

func (f *File) Key(s string) *Token {
	t := f.get()
	if t.T != TTKey {
		f.fail("keyword", s, t)
	}
	if t.V != s {
		f.fail("keyword", s, t)
	}
	return f.ret(t)
}

func (f *File) Ident(s string) *Token {
	t := f.get()
	if t.T != TTWrd {
		f.fail("token", s, t)
	}
	if t.V != s {
		f.fail("token", s, t)
	}
	return f.ret(t)
}

func (f *File) Symbol(s string) *Token {
	t := f.get()
	if t.T != TTSym {
		f.fail("symbol", s, t)
	}
	if t.V != s {
		f.fail("symbol", s, t)
	}
	return f.ret(t)
}

func (f *File) TryV(fn func(string) *Token, s string) (res *Token) {
	res = nil
	defer func() {
		if rcv := recover(); rcv != nil {
			res = nil
			HandleTry(rcv)
		}
	}()
	return fn(s)
}

func (f *File) Try(fn func(string) *Token, s string) (res bool) {
	v := f.TryV(fn, s)
	return v != nil
}
