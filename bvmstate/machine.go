package bvmstate

import (
	"math/big"

	"github.com/benchlab/bvmCreate"
)

func Pop(bench *bvmCreate.BVM) {
	doPop(bench.chainMachine)
}

func doPop(s *bvmCreate.chainMachine) {
	s.Pop()
}

func Push(bench *bvmCreate.BVM) {
	doPush(bench.chainMachine, bench.Input)
}

func doPush(s *bvmCreate.chainMachine, i bvmCreate.Input) {
	size := i.Code().Next(1)
	pushSize := int(new(big.Int).SetBytes(size).Int64())
	s.Push(i.Code().Next(pushSize))
}

func Dup(bench *bvmCreate.BVM) {
	doDup(bench.chainMachine, bench.Input)
}

func doDup(s *bvmCreate.chainMachine, i bvmCreate.Input) {
	size := int(new(big.Int).SetBytes(i.Code().Next(1)).Int64())
	dupSize := int(new(big.Int).SetBytes(i.Code().Next(size)).Int64())
	s.Dup(dupSize)
}

func Swap(bench *bvmCreate.BVM) {
	doSwap(bench.chainMachine, bench.Input)
}

func doSwap(s *bvmCreate.chainMachine, i bvmCreate.Input) {
	size := int(new(big.Int).SetBytes(i.Code().Next(1)).Int64())
	swapSize := int(new(big.Int).SetBytes(i.Code().Next(size)).Int64())
	s.Swap(swapSize)
}
