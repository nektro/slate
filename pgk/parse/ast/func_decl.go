package ast

import (
	"fmt"
	"log"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/nektro/slate/pgk/parse/llvm"
)

type FuncDecl struct {
	Args []*Arg
	Ret  *Ref
	Body *Block
}

func (p *FuncDecl) N() {}

func (p *FuncDecl) Compile(mod *ir.Module, globals VarScope, name string) {
	fnprms := map[string]*ir.Param{}
	params := []*ir.Param{}
	for _, item := range p.Args {
		prm := ir.NewParam("", llvm.GetType(item.Type.V()))
		params = append(params, prm)
		fnprms[item.Name] = prm
	}

	fname := ""
	var ftype types.Type = llvm.Void
	if name == "main" {
		fname = name
		ftype = llvm.I32
	}
	fn := mod.NewFunc(fname, ftype, params...)
	blk := fn.NewBlock("")
	p.Body.Compile(mod, globals, fnprms, blk)
	if ftype == llvm.I32 {
		blk.NewRet(llvm.Int32(0))
	} else {
		blk.NewRet(nil)
	}
	globals[name] = fn
}

func (p *FuncDecl) DependsOn() []string {
	res := []string{}
	for _, item := range p.Body.Calls {
		switch item.(type) {
		case *FuncCall:
			res = append(res, item.(*FuncCall).DependsOn(p.Args)...)
		case *InlineAsm:
			// none
		default:
			log.Fatalln("compile:", "func_decl:", "cannot determine dependency of:", fmt.Sprintf("%#v", item))
		}
	}
	return res
}
