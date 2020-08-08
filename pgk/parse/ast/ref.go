package ast

import (
	"strings"
)

type Ref struct {
	Dots []string
}

func (p *Ref) N() {}

func (p *Ref) V() string { return strings.Join(p.Dots, ".") }
