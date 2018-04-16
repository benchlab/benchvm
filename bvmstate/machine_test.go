package bvmstate

import (
	"math/big"
	"testing"

	"github.com/benchlab/bvmCreate"

	"github.com/benchlab/bvmUtils"
)

func tPop(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	bvmUtils.Assert(t, s.Size() == 2, "BVM Error! chainMachine has a sizing issue.")
	doPop(s)
	bvmUtils.Assert(t, s.Size() == 1, "BVM Error! chainMachine has a sizing issue.")
	doPop(s)
	bvmUtils.Assert(t, s.Size() == 0, "BVM Error! chainMachine has a sizing issue.")
}

func tPush(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	bvmUtils.Assert(t, s.Size() == 0, "BVM Error! chainMachine has a sizing issue.")

}

func tDup(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	bvmUtils.Assert(t, s.Size() == 0, "BVM Error! chainMachine has a sizing issue.")

}

func tSwitch(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	bvmUtils.Assert(t, s.Size() == 0, "BVM Error! chainMachine has a sizing issue.")
}
