package ast

import (
	"log"

	"github.com/llir/llvm/ir"
)

type Node interface {
	N()
}

type Compilable interface {
	Node
	Compile(Node) *ir.Module
}

func Compile(n Node) *ir.Module {
	m := ir.NewModule()

	p, ok := n.(*Program)
	if !ok {
		log.Fatalln("compile failure:", "ast node is not a Program")
	}
	p.Compile(m)

	return m
}
