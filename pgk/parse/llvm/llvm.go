package llvm

import (
	"log"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// globals
var (
	// types
	I8      = types.I8
	I16     = types.I16
	I32     = types.I32
	I64     = types.I64
	I128    = types.I128
	Void    = types.Void
	Ptr     = types.NewPointer
	I8Ptr   = types.I8Ptr
	Func    = types.NewFunc
	StringT = Ptr(types.NewStruct(I8Ptr, I64))

	// const
	Int           = constant.NewInt
	GetElementPtr = constant.NewGetElementPtr
	CharArray     = constant.NewCharArrayFromString
	Array         = constant.NewArray
	PtrToInt      = constant.NewPtrToInt
	IntToPtr      = constant.NewIntToPtr

	// values
	Zero = Int(I64, 0)
	One  = Int(I64, 1)
)

func GetType(s string) types.Type {
	switch s {
	case "Int":
		return I64
	case "String":
		return StringT
	}
	log.Fatalln("compile failure:", "cannot determine type from name:", s)
	return nil
}

func String(mod *ir.Module, in string) *ir.Global {
	s := in + "\x00"
	slen := int64(len(s))
	i := constant.NewCharArrayFromString(s)
	str := mod.NewGlobalDef("", i)
	ptr := constant.NewGetElementPtr(types.NewArray(uint64(slen), I8), str, Int32(0), Int32(0))
	ptr.InBounds = true
	j := constant.NewStruct(StringT.ElemType.(*types.StructType), ptr, Int(I64, slen))
	res := mod.NewGlobalDef("", j)
	return res
}

func Int32(n int64) *constant.Int {
	return Int(I32, n)
}
