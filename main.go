package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/nektro/go-util/util"
	"github.com/nektro/slate/pgk/parse/ast"
	"github.com/nektro/slate/pgk/slate"
)

// Version takes in version string from build_all.sh
var Version = "vMASTER"

func main() {
	rand.Seed(time.Now().UnixNano())

	fRun := flag.String("run", "", "")
	flag.Parse()

	if len(*fRun) == 0 {
		log.Fatalln("-run is empty")
	}

	// log.Println("run:", *fRun)

	n, err := filepath.Abs(*fRun)
	util.DieOnError(err)
	f, err := os.Open(n)
	util.DieOnError(err)
	b, err := ioutil.ReadAll(f)
	util.DieOnError(err)
	s := string(b)
	// fmt.Println(s)

	file := slate.Lexer.LexAll(n, s)
	// log.Println("lex:", "success:", "collected", file.Len(), "tokens")
	file.RemoveComments()
	// file.Print()

	node := slate.Parser.Parse(file)
	// log.Println("parse:", "success:", node)

	module := ast.Compile(node)
	out, _ := os.Create("./out.ll")
	module.WriteTo(out)
}
