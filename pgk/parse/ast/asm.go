package ast

import (
	"log"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
	"github.com/nektro/slate/pgk/parse/llvm"
)

type InlineAsm struct {
	Call       string
	Constraint string
	Args       []*Ref
}

func (p *InlineAsm) N() {}

func (p *InlineAsm) Compile(mod *ir.Module, fnprms map[string]*ir.Param, block *ir.Block) {
	params := []*ir.Param{}
	for _, item := range p.Args {
		n, ok := fnprms[item.V()]
		if !ok {
			log.Fatalln("compile:", "asm:", "unable to find param value for symbol:", item.V())
		}
		params = append(params, n)
	}

	asmFunc := ir.NewInlineAsm(llvm.Ptr(llvm.Func(llvm.I64)), p.Call, p.Constraint)
	asmFunc.SideEffect = true

	params2 := []value.Value{}
	for _, item := range params {
		params2 = append(params2, item)
	}
	block.NewCall(asmFunc, params2...)
}
