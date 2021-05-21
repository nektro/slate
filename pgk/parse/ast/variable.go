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
		if v, ok := p.Value.(*lex.Token); ok {
			if v.T == lex.TTNum {
				g := mod.NewGlobalDef("", constant.NewInt(llvm.GetType("int").(*types.IntType), v.IV()))
				globals[p.Name.Name] = CGValue{g, v}
				return
			}
			if v.T == lex.TTStr {
				g := llvm.String(mod, v.SV())
				globals[p.Name.Name] = CGValue{g, v}
				return
			}
			log.Fatalln("compile failure:", "variable", "unhandled token:", v.T)
		}
		if v, ok := p.Value.(*FuncCall); ok {
			f := globals[v.Name]
			llf, okk := f.LLvmV.(*ir.Func)
			asfd := f.ASTgV.(*FuncDecl)
			if !okk {
				log.Fatalln("compile:", "variable:", "func_call:", "symbol '"+v.Name+"' is not a function")
			}
			if len(llf.Params) > len(v.Args) {
				args := v.Args
				for _, item := range asfd.Args[len(v.Args):] {
					args = append(args, &Ref{[]string{item.Name}})
				}
				d := &FuncDecl{
					Args: asfd.Args[len(v.Args):],
					Ret:  asfd.Ret,
					Body: &Block{[]Node{
						&FuncCall{v.Name, false, args},
					}},
				}
				d.Compile(mod, globals, p.Name.Name)
				return
			}
			log.Fatalln("compile:", "variable:", "func_call:", "only currying is implemented")
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
	case *FuncCall:
		return []string{p.Value.(*FuncCall).Name}
	default:
		log.Fatalln("compile:", "depends:", "variable:", "unhandled type:", p.Value)
	}
	return nil
}
