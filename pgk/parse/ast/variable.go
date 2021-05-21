package ast

import (
	"log"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/nektro/slate/pgk/lex"
	"github.com/nektro/slate/pgk/parse/llvm"
)

type Variable struct {
	Scope
	Name   *Arg
	Public bool
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
		v, ok := p.Value.(*lex.Token)
		if ok {
			if v.T == lex.TTNum {
				g := mod.NewGlobalDef("", constant.NewInt(llvm.GetType("int").(*types.IntType), v.IV()))
				globals[p.Name.Name] = g
				return
			}
			if v.T == lex.TTStr {
				g := llvm.String(mod, v.SV())
				globals[p.Name.Name] = g
				return
			}
			log.Fatalln("compile failure:", "variable", "unhandled token:", v.T)
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
	case *lex.Token:
		return []string{}
	default:
		log.Fatalln("compile:", "depends:", "variable:", "unhandled type:", p.Value)
	}
	return nil
}
