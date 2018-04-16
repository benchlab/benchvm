package bvmstate

import (
	"math/big"
	"testing"

	"github.com/benchlab/bvmUtils"
	"github.com/benchlab/bvmCreate"
)

func tLt(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	doLt(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Error! chainMachine has sizing issues.")
	c := new(big.Int).SetBytes(s.Pop())
	bvmUtils.Assert(t, c.Cmp(new(big.Int).SetUint64(1)) == 0, "BVM Error! Wrong lt value")
}

func tGt(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := 2, 4
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	doGt(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Error! chainMachine has sizing issues.")
	c := new(big.Int).SetBytes(s.Pop())
	bvmUtils.Assert(t, c.Cmp(new(big.Int).SetUint64(1)) == 0, "BVM Error! Wrong gt value")
}

func tSLt(t *testing.T) {

}

func tSGt(t *testing.T) {

}

func tEq(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := 4, 4
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	doEq(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Error! chainMachine has sizing issues")
	c := new(big.Int).SetBytes(s.Pop())
	bvmUtils.Assert(t, c.Cmp(new(big.Int).SetUint64(1)) == 0, "BVM Error! Wrong eq value")
}

func tIsZero(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1 := 0
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	doIsZero(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Error! chainMachine has sizing issues")
	c := new(big.Int).SetBytes(s.Pop())
	bvmUtils.Assert(t, c.Cmp(new(big.Int).SetUint64(1)) == 0, "BVM Error! Wrong isZero value")
}
