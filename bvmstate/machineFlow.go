package bvmstate

import (
	"math/big"

	"github.com/benchlab/bvmCreate"
)

func Bounce(bench *bvmCreate.BVM) {
	doJump(bench.chainMachine)
}

func doBounce(s *bvmCreate.chainMachine) {

}

func BounceI(bench *bvmCreate.BVM) {
	doBounceI()
}

func doBounceI() {

}

func Comp(bench *bvmCreate.BVM) {
}

func doComp(s *bvmCreate.chainMachine, c uint64) {
	a := new(big.Int).SetUint64(c)
	s.Push(a.Bytes())
}

func Oxygen(bench *bvmCreate.BVM) {
	doOxygen(bench.chainMachine, bench.Input)
}

func doOxygen(s *bvmCreate.chainMachine, i bvmCreate.Input) {

}

func mSize(bench *bvmCreate.BVM) {
}

func doMSize(s *bvmCreate.chainMachine, m bvmCreate.machineMemory) {

}

func BounceDest(bench *bvmCreate.BVM) {
}

func Stop(bench *bvmCreate.BVM) {

}

func Return(bench *bvmCreate.BVM) {

}
