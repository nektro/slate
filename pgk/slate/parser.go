package slate

import (
	"fmt"
	"reflect"

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

func init() {
	lex.Errors = append(
		lex.Errors,
		"parse failure",
	)
}

func (p *ParserType) fail(s string) {
	panic(fmt.Sprint("parse failure: ", s))
}

//
//

func (p *ParserType) Parse(f *lex.File) ast.Node {
	return p.Program(f)
}

func (p *ParserType) Try(f *lex.File, fn func(*lex.File) ast.Node) (_ ast.Node, res bool) {
	defer func() {
		if rcv := recover(); rcv != nil {
			res = lex.HandleTry(rcv)
		}
	}()
	return fn(f), true
}

func (p *ParserType) TryV(f *lex.File, fn func(*lex.File) ast.Node) ast.Node {
	n, _ := p.Try(f, fn)
	return n
}

func (p *ParserType) Program(f *lex.File) ast.Node {
	res := new(ast.Program)
	for {
		public := f.Try(f.Key, "pub")

		if f.Try(f.Key, "const") {
			res.Stmts = append(res.Stmts, p.Variable(f, public, true).(*ast.Variable))
			continue
		}
		break
	}
	return res
}

func (p *ParserType) Variable(f *lex.File, pub, cnst bool) ast.Node {
	n := p.Arg(f).(*ast.Arg)
	f.Symbol("=")
	defer f.Symbol(";")

	if f.Try(f.Key, "func") {
		return &ast.Variable{ast.ScopeGlobal, n, pub, cnst, p.FuncDecl(f)}
	}
	if x := f.TryV(f.Type, lex.TTNum.S()); x != nil {
		return &ast.Variable{ast.ScopeGlobal, n, pub, cnst, x}
	}

	p.fail("variable: didn't see func, saw " + f.Get().String())
	return nil
}

func (p *ParserType) FuncDecl(f *lex.File) ast.Node {
	res := new(ast.FuncDecl)

	f.Symbol("(")
	for {
		n, ok := p.Try(f, p.Arg)
		if !ok {
			break
		}
		res.Args = append(res.Args, n.(*ast.Arg))
	}
	f.Symbol(")")

	if f.Try(f.Symbol, ":") {
		res.Ret = p.Ref(f).(*ast.Ref)
	}

	res.Body = p.Block(f).(*ast.Block)

	return res
}

func (p *ParserType) Arg(f *lex.File) ast.Node {
	defer f.Try(f.Symbol, ",")
	t1 := f.Type(lex.TTWrd.S())
	if f.Try(f.Symbol, ":") {
		t2 := p.Ref(f)
		return &ast.Arg{t1.V, t2.(*ast.Ref)}
	}
	return &ast.Arg{t1.V, &ast.Ref{}}
}

func (p *ParserType) Ref(f *lex.File) ast.Node {
	dots := []string{}
	dots = append(dots, f.Type(lex.TTWrd.S()).V)
	for f.Try(f.Symbol, ".") {
		dots = append(dots, f.Type(lex.TTWrd.S()).V)
	}
	return &ast.Ref{dots}
}

func (p *ParserType) Block(f *lex.File) ast.Node {
	res := new(ast.Block)
	f.Symbol("{")
	for true {
		n, ok := p.Try(f, p.FuncCall)
		if ok {
			res.Calls = append(res.Calls, n.(*ast.FuncCall))
			continue
		}
		n, ok = p.Try(f, p.InlineAsm)
		if ok {
			res.Calls = append(res.Calls, n.(*ast.InlineAsm))
			continue
		}
		if f.Try(f.Symbol, "}") {
			break
		}
	}
	return res
}

func (p *ParserType) FuncCall(f *lex.File) ast.Node {
	addParam := func(f *lex.File, c *ast.FuncCall, n ast.Node) {
		rv := reflect.ValueOf(n)
		if !(rv.Kind() == reflect.Kind(0)) && !rv.IsNil() {
			c.Params = append(c.Params, n)
			f.Symbol(",")
		}
	}
	res := new(ast.FuncCall)
	res.Builtin = f.Try(f.Symbol, "@")
	res.Name = f.Type(lex.TTWrd.S()).V
	f.Symbol("(")
	l := len(res.Params)
	for {
		addParam(f, res, f.TryV(f.TypeS, lex.TTStr.S()))
		addParam(f, res, f.TryV(f.TypeS, lex.TTNum.S()))
		addParam(f, res, p.TryV(f, p.Ref))
		//
		if f.Try(f.Symbol, ")") {
			break
		}
		ll := len(res.Params)
		if ll != l {
			l = ll
			continue
		}
		p.fail("unable to parse function parameters")
	}
	return res
}

func (p *ParserType) InlineAsm(f *lex.File) ast.Node {
	res := &ast.InlineAsm{}
	f.Key("asm")
	res.Call = f.Type(lex.TTStr.S()).SV()
	f.Symbol(",")
	res.Constraint = f.Type(lex.TTStr.S()).SV()
	for {
		if !f.Try(f.Symbol, ",") {
			break
		}
		res.Args = append(res.Args, p.Ref(f).(*ast.Ref))
	}
	return res
}
