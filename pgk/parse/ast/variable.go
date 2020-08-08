package ast

import (
	"log"

	"github.com/llir/llvm/ir"
)

type Variable struct {
	Scope
	Name   *Arg
	Public bool
	Const  bool
	Value  Node
}

type Scope int

const (
	ScopeGlobal Scope = iota
	ScopeFunc
	ScopeStruct
)

func (p *Variable) N() {}

func (p *Variable) Compile(mod *ir.Module, globals VarScope) {
	switch p.Scope {
	case ScopeGlobal:
		f, ok := p.Value.(*FuncDecl)
		if ok {
			f.Compile(mod, globals, p.Name.Name)
			return
		}
		log.Fatalln("compile failure:", "variable:", "unhandled type:", p.Value)
	default:
		log.Fatalln("compile failure:", "variable:", "unhandled scope:", p.Scope)
	}
}

func (p *Variable) DependsOn() []string {
	switch p.Value.(type) {
	case *FuncDecl:
		return p.Value.(*FuncDecl).DependsOn()
	default:
		log.Fatalln("compile:", "depends:", "variable:", "unhandled type:", p.Value)
	}
	return nil
}
