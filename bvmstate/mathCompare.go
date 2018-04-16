package bvmstate

import (
	"math/big"

	"github.com/benchlab/bvmCreate"
)

func Lt(bench *bvmCreate.BVM) {
	doLt(bench.chainMachine)
}

func doLt(s *bvmCreate.chainMachine) {
	x := bigInt(s.Pop())
	y := bigInt(s.Pop())
	if x.Cmp(y) < 0 {
		s.Push(new(big.Int).SetUint64(1).Bytes())
	} else {
		s.Push(new(big.Int).Bytes())
	}
}

func Gt(bench *bvmCreate.BVM) {
	doGt(bench.chainMachine)
}

func doGt(s *bvmCreate.chainMachine) {
	x := bigInt(s.Pop())
	y := bigInt(s.Pop())
	if x.Cmp(y) > 0 {
		s.Push(new(big.Int).SetUint64(1).Bytes())
	} else {
		s.Push(new(big.Int).Bytes())
	}
}

func SLt(bench *bvmCreate.BVM) {

}

func SGt(bench *bvmCreate.BVM) {

}

func Eq(bench *bvmCreate.BVM) {
	doEq(bench.chainMachine)
}

func doEq(s *bvmCreate.chainMachine) {
	x := bigInt(s.Pop())
	y := bigInt(s.Pop())
	if x.Cmp(y) == 0 {

		s.Push(new(big.Int).SetUint64(1).Bytes())
	} else {
		s.Push(new(big.Int).Bytes())
	}
}

func IsZero(bench *bvmCreate.BVM) {
	doIsZero(bench.chainMachine)
}

func doIsZero(s *bvmCreate.chainMachine) {
	x := bigInt(s.Pop())
	if x.Sign() > 0 {
		s.Push(new(big.Int).Bytes())
	} else {
		s.Push(new(big.Int).SetUint64(1).Bytes())
	}
}
