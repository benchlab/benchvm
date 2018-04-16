package bvmstate

import (
	"math/big"

	"github.com/benchlab/bvm/memory"
	"github.com/benchlab/bvmCreate"
)

func MemSize(bench *bvmCreate.BVM) {
}

func doMemSize(s *bvmCreate.chainMachine, m bvmCreate.machineMemory) {
	fm := m.(memory.machineMemory)
	size := new(big.Int).SetInt64(int64(fm.Size()))
	s.Push(size.Bytes())
}

func Get(bench *bvmCreate.BVM) {
}

func doGet(s *bvmCreate.chainMachine, m bvmCreate.machineMemory) {
	offset := bigInt(s.Pop())
	fm := m.(memory.machineMemory)
	val := fm.Get(offset.Int64(), 32)
	s.Push(val)
}

func Set(bench *bvmCreate.BVM) {
}

func doSet(s *bvmCreate.chainMachine, m bvmCreate.machineMemory) {
	start := s.Pop()
	val := s.Pop()
	fm := m.(memory.machineMemory)
	fm.Set(new(big.Int).SetBytes(start).Uint64(), 32, val)
}

func Load(bench *bvmCreate.BVM) {
	doLoad(bench.chainMachine, bench.Input, bench.State)
}

func doLoad(s *bvmCreate.chainMachine, i bvmCreate.Input, bvmstate bvmCreate.machineState) {

}

func Store(bench *bvmCreate.BVM) {
	doStore(bench.chainMachine, bench.Input, bench.State)
}

func doStore(s *bvmCreate.chainMachine, i bvmCreate.Input, bvmstate bvmCreate.machineState) {

}
