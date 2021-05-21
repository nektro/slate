package ast

import (
	"fmt"
	"log"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
	"github.com/nektro/go-util/arrays/stringsu"
	"github.com/nektro/slate/pgk/lex"
	"github.com/nektro/slate/pgk/parse/llvm"
)

type FuncCall struct {
	Name    string
	Builtin bool
	Args    []Node
}

func (p *FuncCall) N() {}

func (p *FuncCall) Compile(mod *ir.Module, globals VarScope, fnprms map[string]*ir.Param, block *ir.Block) {
	params := []value.Value{}
	for _, item := range p.Args {
		switch item.(type) {
		case *lex.Token:
			v := item.(*lex.Token)
			switch v.T {
			case lex.TTStr:
				params = append(params, llvm.String(mod, v.SV()))
			case lex.TTNum:
				params = append(params, llvm.Int(llvm.I64, v.IV()))
			default:
				log.Fatalln("compile:", "func_call:", "can't determine param value for token type:", v.PL())
			}
		case *Ref:
			ref := item.(*Ref)
			name := ref.V()
			if !strings.ContainsRune(name, '.') {
				if pv, ok := fnprms[name]; ok {
					params = append(params, pv)
					continue
				}
				if gv, ok := globals[name]; ok {
					params = append(params, gv.LLvmV)
					continue
				}
				log.Fatalln("compile:", "func_call:", "can't determine arg value for", name)
			}
			pr, ok := fnprms[ref.Dots[0]]
			if !ok {
				log.Fatalln("compile:", "func_call:", "can't determine param value for", name)
			}
			if pr.Type() == llvm.StringT {
				if ref.Dots[1] == "ptr" {
					gep := block.NewGetElementPtr(llvm.StringT.ElemType, pr, llvm.Int32(0), llvm.Int32(0))
					gep.InBounds = true
					val := block.NewLoad(llvm.I8Ptr, gep)
					params = append(params, block.NewPtrToInt(val, llvm.I64))
					continue
				}
				if ref.Dots[1] == "len" {
					gep := block.NewGetElementPtr(llvm.StringT.ElemType, pr, llvm.Int32(0), llvm.Int32(1))
					gep.InBounds = true
					val := block.NewLoad(llvm.I64, gep)
					params = append(params, val)
					continue
				}
			}
			log.Fatalln("compile:", "func_call:", "can't determine param value for", name)
		default:
			log.Fatalln("compile:", "func_call:", "can't determine param value for", fmt.Sprintf("%#v", item))
		}
	}
	if p.Builtin {
		log.Fatalln("compile:", "func_call:", "unhandled builtin:", p.Name)
		return
	}
	callee, ok := globals[p.Name]
	if !ok {
		log.Fatalln("compile:", "func_call:", "can't find symbol", p.Name)
	}
	block.NewCall(callee.LLvmV, params...)
}

func (p *FuncCall) DependsOn(fnparams []*Arg) []string {
	res := []string{p.Name}
	paramnames := []string{}
	for _, item := range fnparams {
		paramnames = append(paramnames, item.Name)
	}
	argnames := []string{}
	for _, item := range p.Args {
		switch item.(type) {
		case *Ref:
			v := item.(*Ref)
			argnames = append(argnames, v.Dots[0])
		}
	}
	for _, item := range argnames {
		if !stringsu.Contains(paramnames, item) {
			res = append(res, item)
		}
	}
	return res
}
