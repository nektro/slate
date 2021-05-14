package lex

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/nektro/go-util/ansi/style"
)

// Token is a
type Token struct {
	T TokenType // type
	V string    // value
	L int       // line
	P int       // pos
	S int       // size (len(V))
}

func (p *Token) N() {}

func (v *Token) String() string {
	return fmt.Sprintf("%s{%s:%d}", v.PL(), v.LP(), v.S)
}

// PL returns T_type(value)
func (v *Token) PL() string {
	return fmt.Sprintf("T_%s("+style.BgRed+"%s"+style.ResetBgColor+")", v.T, v.V)
}

// LP returns line:pos
func (v *Token) LP() string {
	return fmt.Sprintf("%d:%d", v.L, v.P)
}

func (v *Token) SV() string {
	if v.T != TTStr {
		log.Fatalln("lex:", "token:", "called SV() on non-string token")
	}
	s := v.V
	s = s[1 : len(s)-1]
	s = strings.ReplaceAll(s, "\\n", "\n")
	for {
		i := strings.IndexRune(s, '\\')
		if i == -1 {
			break
		}
		switch s[i+1] {
		case 'x':
			x, _ := strconv.ParseInt(s[i+2:i+4], 16, 64)
			s = strings.Replace(s, s[i:i+4], string(x), 1)
		}
	}
	return s
}

func (v *Token) IV() int64 {
	if v.T != TTNum {
		log.Fatalln("lex:", "token:", "called IV() on non-number token")
	}
	n, _ := strconv.ParseInt(v.V, 10, 64)
	return n
}

//
//

type TokenType string

// TokenType enum
const (
	TTWrd TokenType = "W"
	TTCom TokenType = "C"
	TTStr TokenType = "Q"
	TTKey TokenType = "K"
	TTSym TokenType = "S"
	TTNum TokenType = "N"
)

func (tt TokenType) S() string {
	return string(tt)
}
