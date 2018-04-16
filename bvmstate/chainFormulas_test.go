package bvmstate

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/benchlab/bvmUtils"
	"github.com/benchlab/bvmCreate"
)

func tPlusPos(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := 4, 2
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	doAddition(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! Stack is the wrong size")
	c := new(big.Int).SetBytes(s.Pop())
	bvmUtils.Assert(t, c.Cmp(intAsBig(o1+o2)) == 0, "BVM Formula Error! Addition formula has the wrong value.")
}

func tPlusNeg(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := -4, -2
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	doAddition(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! chainMachine is the wrong size")
	c := new(big.Int).SetBytes(s.Pop())
	bvmUtils.Assert(t, c.Cmp(intAsBig(o1+o2)) == 0, "BVM Formula Error! Addition formula has the wrong value.")
}

func tPlusZero(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := -4, 0
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	doAddition(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! chainMachine is the wrong size")
	c := new(big.Int).SetBytes(s.Pop())
	bvmUtils.Assert(t, c.Cmp(intAsBig(o1+o2)) == 0, "BVM Formula Error! Addition formula has the wrong value.")
}

func tMinusPos(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := 4, 2
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	xSub(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! chainMachine is the wrong size")
	c := new(big.Int).SetBytes(s.Pop())
	bvmUtils.Assert(t, c.Cmp(intAsBig(o1-o2)) == 0, "BVM Formula Error! Subtraction formula has the wrong value.")
}

func tMinusNeg(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := -4, -2
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	xSub(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! chainMachine is the wrong size")
	c := new(big.Int).SetBytes(s.Pop())
	bvmUtils.Assert(t, c.Cmp(intAsBig(o1-o2)) == 0, "BVM Formula Error! Subtraction formula has the wrong value.")
}

func tMinusZero(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := -4, 0
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	xSub(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! chainMachine is the wrong size")
	c := new(big.Int).SetBytes(s.Pop())
	bvmUtils.Assert(t, c.Cmp(intAsBig(o1-o2)) == 0, "BVM Formula Error! Subtraction formula has the wrong value.")
}

func tMulti(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := 4, 2
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	xMult(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! chainMachine is the wrong size")
	c := new(big.Int).SetBytes(s.Pop())
	bvmUtils.Assert(t, c.Cmp(intAsBig(o1*o2)) == 0, "BVM Formula Error! Multiplication formula has the wrong value.")
}

func tDivision(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := 4, 2
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	doDivision(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! chainMachine is the wrong size")
	c := new(big.Int).SetBytes(s.Pop())
	bvmUtils.Assert(t, c.Cmp(intAsBig(o1/o2)) == 0, fmt.Sprintf("BVM Formula Error! Division formula has the wrong value: %d", c.Int64()))
}

func tMod(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := 4, 2
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	doModulo(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! chainMachine is the wrong size")
	c := new(big.Int).SetBytes(s.Pop())
	bvmUtils.Assert(t, c.Cmp(intAsBig(o1%o2)) == 0, fmt.Sprintf("BVM Formula Error! Module has the wrong value:: %d", c.Int64()))
}

func tPlusMod(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2, o3 := 4, 2, 1
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	s.Push(intAsBig(o3).Bytes())
	doAddMod(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! chainMachine is the wrong size")

}

func tMultiMod(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2, o3 := 4, 2, 1
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	s.Push(intAsBig(o3).Bytes())
	doMulMod(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! chainMachine is the wrong size")
}

func tConBF(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	s1 := "hola"
	s2 := "universe"
	s.Push([]byte(s1))
	s.Push([]byte(s2))
	doConcat(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! chainMachine is the wrong size")
	p := s1 + s2
	bvmUtils.Assert(t, len(s.Pop()) == len([]byte(p)), "BVM Formula Error! Wrong Value Used.")
}

func tConBE(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	s1 := ""
	s2 := ""
	s.Push([]byte(s1))
	s.Push([]byte(s2))
	doConcat(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! chainMachine is the wrong size")
	p := s1 + s2
	bvmUtils.Assert(t, len(s.Pop()) == len([]byte(p)), "BVM Formula Error! Wrong Value Used.")
}

func tConOE(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	s1 := "hello"
	s2 := ""
	s.Push([]byte(s1))
	s.Push([]byte(s2))
	doConcat(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! chainMachine is the wrong size")
	p := s1 + s2
	bvmUtils.Assert(t, len(s.Pop()) == len([]byte(p)), "BVM Formula Error! Wrong Value Used.")
}
