package slate

import (
	"github.com/nektro/slate/pgk/lex"
	"github.com/nektro/slate/pgk/parse/ast"
)

// singleton
var (
	Parser = new(ParserType)
)

// ParserType is a
type ParserType struct {
	//
}

//
//

// func init() {
// 	lex.Errors = append(
// 		lex.Errors,
// 		"parse failure",
// 	)
// }

// func (p *ParserType) fail(s string) {
// 	panic(fmt.Sprint("parse failure: ", s))
// }

// //
// //

func (p *ParserType) Parse(f *lex.File) ast.Node {
	return p.Program(f)
}

// func (p *ParserType) Try(f *lex.File, fn func(*lex.File) ast.Node) (_ ast.Node, res bool) {
// 	defer func() {
// 		if rcv := recover(); rcv != nil {
// 			res = lex.HandleTry(rcv)
// 		}
// 	}()
// 	return fn(f), true
// }

// func (p *ParserType) TryV(f *lex.File, fn func(*lex.File) ast.Node) ast.Node {
// 	n, _ := p.Try(f, fn)
// 	return n
// }

func (p *ParserType) Program(f *lex.File) ast.Node {
	res := new(ast.Program)
	for !f.Done() {
		public := f.Try(f.Key, "pub")
		f.Key("const")

		res.Stmts = append(res.Stmts, p.Variable(f, public, true).(*ast.Variable))
		f.Symbol(";")
	}
	return res
}

func (p *ParserType) Variable(f *lex.File, pub bool, decl bool) ast.Node {
	n := p.Arg(f).(*ast.Arg)
	if decl {
		f.Symbol("=")
		v := p.Value(f)
		if v == nil {
			panic("slate: parse: variable: 2")
		}
		return &ast.Variable{ast.ScopeGlobal, n, pub, v}
	}
	panic("slate: parse: variable: 1")
}

func (p *ParserType) Arg(f *lex.File) ast.Node {
	t1 := f.Type(lex.TTWrd)
	if f.Try(f.Symbol, ":") {
		t2 := p.Ref(f, nil)
		return &ast.Arg{t1.V, t2.(*ast.Ref)}
	}
	return &ast.Arg{t1.V, nil}
}

func (p *ParserType) Value(f *lex.File) ast.Node {
	if f.Try(f.Key, "func") {
		return p.FuncDecl(f)
	}
	if x := f.TryV(f.TypeS, string(lex.TTNum)); x != nil {
		return x
	}
	if x := f.TryV(f.TypeS, string(lex.TTStr)); x != nil {
		return x
	}
	if x := f.TryV(f.TypeS, string(lex.TTWrd)); x != nil {
		if f.Try(f.Symbol, "(") {
			f.Back(1)
			return p.FuncCall(f, x.V, false)
		}
		return p.Ref(f, &x.V)
	}
	return nil
}

func (p *ParserType) FuncDecl(f *lex.File) ast.Node {
	res := new(ast.FuncDecl)

	f.Symbol("(")
	if !f.Try(f.Symbol, ")") {
		res.Args = append(res.Args, p.Arg(f).(*ast.Arg))
		for {
			if f.Try(f.Symbol, ")") {
				break
			}
			f.Symbol(",")
			res.Args = append(res.Args, p.Arg(f).(*ast.Arg))
		}
	}
	res.Body = p.Block(f).(*ast.Block)
	return res
}

func (p *ParserType) Ref(f *lex.File, first *string) ast.Node {
	dots := []string{}
	if first != nil {
		dots = append(dots, *first)
	} else {
		dots = append(dots, f.Type(lex.TTWrd).V)
	}
	for f.Try(f.Symbol, ".") {
		dots = append(dots, f.Type(lex.TTWrd).V)
	}
	return &ast.Ref{dots}
}

func (p *ParserType) Block(f *lex.File) ast.Node {
	res := new(ast.Block)
	f.Symbol("{")
	for {
		n := p.Stmt(f)
		if n == nil {
			break
		}
		res.Calls = append(res.Calls, n)
		f.Symbol(";")
	}
	f.Symbol("}")
	return res
}

func (p *ParserType) Stmt(f *lex.File) ast.Node {
	if f.Try(f.Symbol, "@") {
		n := f.Type(lex.TTWrd)
		return p.FuncCall(f, n.V, true)
	}
	if n := f.TryV(f.TypeS, string(lex.TTWrd)); n != nil {
		return p.FuncCall(f, n.V, false)
	}
	if f.Try(f.Key, "asm") {
		return p.InlineAsm(f)
	}
	return nil
}

func (p *ParserType) FuncCall(f *lex.File, name string, builtin bool) ast.Node {
	args := []ast.Node{}
	f.Symbol("(")
	for {
		n := p.Value(f)
		if n == nil {
			f.Symbol(")")
			break
		}
		args = append(args, n)
		if f.Try(f.Symbol, ")") {
			break
		}
		f.Symbol(",")
	}
	return &ast.FuncCall{name, builtin, args}
}

func (p *ParserType) InlineAsm(f *lex.File) ast.Node {
	res := new(ast.InlineAsm)
	res.Call = f.Type(lex.TTStr).SV()
	f.Symbol(",")
	res.Constraint = f.Type(lex.TTStr).SV()
	for {
		if !f.Try(f.Symbol, ",") {
			break
		}
		res.Args = append(res.Args, p.Ref(f, nil).(*ast.Ref))
	}
	return res
}
