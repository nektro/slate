package ast

import (
	"log"
	"reflect"

	"github.com/llir/llvm/ir"
)

type Block struct {
	Calls []Node
}

func (p *Block) N() {}

func (p *Block) Compile(mod *ir.Module, globals VarScope, fnprms map[string]*ir.Param, block *ir.Block) {
	for _, item := range p.Calls {
		switch item.(type) {
		case *FuncCall:
			item.(*FuncCall).Compile(mod, globals, fnprms, block)
		case *InlineAsm:
			item.(*InlineAsm).Compile(mod, fnprms, block)
		default:
			log.Fatalln("compile:", "block:", "invalid statement node type:", reflect.TypeOf(item))
		}
	}
}
