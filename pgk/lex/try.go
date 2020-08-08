package lex

import (
	"fmt"
	"strings"
)

var Errors = []string{
	"expected token",
	"expected token type",
}

func HandleTry(rcv interface{}) bool {
	rcvs := fmt.Sprintf("%v", rcv)
	// log.Println("debug:", rcvs)
	for _, item := range Errors {
		if strings.TrimPrefix(rcvs, item+":") != rcvs {
			return false
		}
	}
	panic(rcvs)
}
