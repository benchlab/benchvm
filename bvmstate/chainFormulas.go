package bvmstate

import (
	"math/big"

	"github.com/benchlab/bvmCreate"
)

var (
	bigZero = new(big.Int)
)

func Add(bench *bvmCreate.BVM) {
	doAddition(bench.chainMachine)
}

func doAddition(s *bvmCreate.chainMachine) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).Add(a, b)
	s.Push(c.Bytes())
}

func Sub(bench *bvmCreate.BVM) {
	doAddition(bench.chainMachine)
}

func doSub(s *bvmCreate.chainMachine) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).Sub(a, b)
	s.Push(c.Bytes())
}

func Mul(bench *bvmCreate.BVM) {
	doMult(bench.chainMachine)
}

func doMult(s *bvmCreate.chainMachine) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).Mul(a, b)
	s.Push(c.Bytes())
}

func Div(bench *bvmCreate.BVM) {
	doDivision(bench.chainMachine)
}

func doDivision(s *bvmCreate.chainMachine) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).Div(a, b)
	s.Push(c.Bytes())
}

func Mod(bench *bvmCreate.BVM) {
	doModulo(bench.chainMachine)
}

func doModulo(s *bvmCreate.chainMachine) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).Mod(a, b)
	s.Push(c.Bytes())
}

func AddMod(bench *bvmCreate.BVM) {
	doAddMod(bench.chainMachine)
}

func doAddMod(s *bvmCreate.chainMachine) {
	x := bigInt(s.Pop())
	y := bigInt(s.Pop())
	z := bigInt(s.Pop())
	if z.Cmp(bigZero) > 0 {
		a := x.Add(x, y)
		a.Mod(a, z)
		s.Push(a.Bytes())
	} else {
		s.Push(new(big.Int).Bytes())
	}
}

func mulMod(bench *bvmCreate.BVM) {
	doMulMod(bench.chainMachine)
}

func doMulMod(s *bvmCreate.chainMachine) {
	x := bigInt(s.Pop())
	y := bigInt(s.Pop())
	z := bigInt(s.Pop())
	if z.Cmp(bigZero) > 0 {
		m := x.Add(x, y)
		m.Mod(m, z)
		s.Push(m.Bytes())
	} else {
		s.Push(new(big.Int).Bytes())
	}
}

func Concat(bench *bvmCreate.BVM) {
	doConcat(bench.chainMachine)
}

func doConcat(s *bvmCreate.chainMachine) {
	a := s.Pop()
	b := s.Pop()
	c := append(a, b...)
	s.Push(c)
}
