package bvmstate

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/benchlab/bvmCreate"

	"github.com/benchlab/bvmUtils"
)

func tAnd(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	doAnd(s)
	bvmUtils.Assert(t, s.Size() == 1, fmt.Sprintf("BVM Formula Error! Wrong chainMachine size: %d", s.Size()))
}

func tOr(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	doOr(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! Wrong chainMachine size.")
}

func tXor(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	doXor(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! Wrong chainMachine size:")
}

func tNot(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1 := 0
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	doNot(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Formula Error! Wrong chainMachine size:")
}
