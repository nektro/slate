package ast

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/nektro/go-util/arrays/stringsu"
)

type Program struct {
	Stmts []*Variable
}

func (p *Program) N() {}

type VarScope map[string]constant.Constant

func (p *Program) Compile(mod *ir.Module) {
	globals := VarScope{}
	available := []string{}
	notdone := []string{}
	for _, item := range p.Stmts {
		notdone = append(notdone, item.Name.Name)
	}
	// fmt.Println("compile:", "notdone:", notdone)
	for {
		if len(notdone) == 0 {
			break
		}
		for _, item := range notdone {
			stmt := getStmt(p.Stmts, item)
			if containsAll(available, stmt.DependsOn()) {
				stmt.Compile(mod, globals)
				notdone = stringsu.Remove(notdone, item)
				available = append(available, item)
				// fmt.Println("compile:", "complete:", item)
				break
			}
		}
	}
}

func containsAll(haystack, needle []string) bool {
	for _, item := range needle {
		if !stringsu.Contains(haystack, item) {
			return false
		}
	}
	return true
}

func getStmt(haystack []*Variable, name string) *Variable {
	for _, item := range haystack {
		if item.Name.Name == name {
			return item
		}
	}
	return nil
}
