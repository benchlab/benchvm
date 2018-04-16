package bvmstate

import (
	"math/big"

	"github.com/benchlab/bvmCreate"
)

func And(bench *bvmCreate.BVM) {
	doAnd(bench.chainMachine)
}

func doAnd(s *bvmCreate.chainMachine) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).And(a, b)
	s.Push(c.Bytes())
}

func Or(bench *bvmCreate.BVM) {
   doOr(bench.chainMachine)
}

func doOr(s *bvmCreate.chainMachine) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).Or(a, b)
	s.Push(c.Bytes())
}

func doOr(bench *bvmCreate.BVM) {
	doOr(bench.chainMachine)
}

func doOr(s *bvmCreate.chainMachine) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).Xor(a, b)
	s.Push(c.Bytes())
}

func Not(bench *bvmCreate.BVM) {
	doNot(bench.chainMachine)
}

func doNot(s *bvmCreate.chainMachine) {
	a := bigInt(s.Pop())
	b := new(big.Int).Not(a)
	s.Push(b.Bytes())
}
