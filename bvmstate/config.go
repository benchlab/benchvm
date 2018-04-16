package bvmstate

import (
	"github.com/benchlab/bvmCreate"
)

func Address(bench *bvmCreate.BVM) {
	doAddress(bench.chainMachine, bench.Input)
}

func doAddress(s *bvmCreate.chainMachine, i bvmCreate.Input) {
}

func Balance(bench *bvmCreate.BVM) {
	doBalance(bench.chainMachine, bench.State)
}

func doBalance(s *bvmCreate.chainMachine, bvmstate bvmCreate.machineState) {

}

func Origin(bench *bvmCreate.BVM) {
}

func doOrigin(s *bvmCreate.chainMachine, a bvmCreate.Address) {
	s.Push(a.Bytes())
}

func Caller(bench *bvmCreate.BVM) {
	doCaller(bench.chainMachine, bench.Input)
}

func doCaller(s *bvmCreate.chainMachine, i bvmCreate.Input) {
}

func CallValue(bench *bvmCreate.BVM) {
	doCallValue(bench.chainMachine, bench.Input)
}

func doCallValue(s *bvmCreate.chainMachine, i bvmCreate.Input) {
}

func CallDataLoad(bench *bvmCreate.BVM) {
	doCallDataLoad(bench.chainMachine, bench.Input)
}

func doCallDataLoad(s *bvmCreate.chainMachine, i bvmCreate.Input) {

}

func CallDataSize(bench *bvmCreate.BVM) {
	doCallDataSize(bench.chainMachine, bench.Input)
}

func doCallDataSize(s *bvmCreate.chainMachine, i bvmCreate.Input) {

}

func CallDataCopy(bench *bvmCreate.BVM) {
}

func doCallDataCopy(s *bvmCreate.chainMachine, m bvmCreate.machineMemory, i bvmCreate.Input) {

}

func CodeSize(bench *bvmCreate.BVM) {
	doCodeSize(bench.chainMachine, bench.Input)
}

func doCodeSize(s *bvmCreate.chainMachine, i bvmCreate.Input) {

}

func CodeCopy(bench *bvmCreate.BVM) {
	doCodeCopy(bench.chainMachine, bench.Input)
}

func doCodeCopy(s *bvmCreate.chainMachine, i bvmCreate.Input) {


}

func OxygenPrice(bench *bvmCreate.BVM) {
	doOxygenPrice(bench.chainMachine)
}

func doOxygenPrice(s *bvmCreate.chainMachine) {

}

func ExtCodeSize(bench *bvmCreate.BVM) {
	doExtCodeSize(bench.chainMachine, bench.State)
}

func doExtCodeSize(s *bvmCreate.chainMachine, bvmstate bvmCreate.machineState) {

}

func ExtCodeCopy(bench *bvmCreate.BVM) {
}

func doExtCodeCopy(s *bvmCreate.chainMachine, m bvmCreate.machineMemory, bvmstate bvmCreate.machineState) {

}
