package ast

type Arg struct {
	Name string
	Type *Ref
}

func (p *Arg) N() {}
