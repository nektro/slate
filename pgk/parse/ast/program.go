package ast

import (
	"log"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/nektro/go-util/arrays/stringsu"
)

type Program struct {
	Stmts []*Variable
}

func (p *Program) N() {}

type CGValue struct {
	LLvmV constant.Constant
	ASTgV Node
}

type VarScope map[string]CGValue

func (p *Program) Compile(mod *ir.Module) {
	ensureNoShadowing(p.Stmts)
	globals := VarScope{}
	available := []string{}
	notdone := stringsu.Depupe(collectSymbolDeps(p.Stmts, &[]string{}, "main"))

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
				break
			}
		}
	}
}

func ensureNoShadowing(haystack []*Variable) {
	names := []string{}
	for _, item := range haystack {
		if stringsu.Contains(names, item.Name.Name) {
			log.Fatalln("compile:", "found shadow of symbol", item.Name.Name)
		}
		names = append(names, item.Name.Name)
	}
}

func collectSymbolDeps(haystack []*Variable, notdone *[]string, name string) []string {
	st := getStmt(haystack, name)
	if st == nil {
		log.Fatalln("compile:", "can't find symbol:", name)
	}
	if name == "main" {
		if !st.Public {
			log.Fatalln("compile:", "main needs to be public")
		}
		if _, ok := st.Value.(*FuncDecl); !ok {
			log.Fatalln("compile:", "main must be a function")
		}
	}
	*notdone = append(*notdone, name)
	for _, item := range st.DependsOn() {
		collectSymbolDeps(haystack, notdone, item)
	}
	return *notdone
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
