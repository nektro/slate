package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/nektro/slate/pgk/parse/ast"
	"github.com/nektro/slate/pgk/slate"
)

// Version takes in version string from build_all.sh
var Version = "vMASTER"

// shorthands
var (
	i32  = types.I32
	zero = constant.NewInt(i32, 0)
	one  = constant.NewInt(i32, 1)
	i8   = types.I8
	Ptr  = types.NewPointer
	i64  = types.I64
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fRun := flag.String("run", "", "")
	flag.Parse()

	if len(*fRun) == 0 {
		log.Fatalln("-run is empty")
	}

	log.Println("run:", *fRun)

	n, _ := filepath.Abs(*fRun)
	f, _ := os.Open(n)
	b, _ := ioutil.ReadAll(f)
	s := string(b)
	// fmt.Println(s)

	file := slate.Lexer.LexAll(n, s)
	log.Println("lex:", "success:", "collected", file.Len(), "tokens")
	file.RemoveComments()
	// file.Print()

	node := slate.Parser.Parse(file)
	log.Println("parse:", "success:", node)

	module := ast.Compile(node)
	out, _ := os.Create("./out.ll")
	module.WriteTo(out)

}
