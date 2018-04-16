package bvmstate

import "github.com/benchlab/bvmCreate"

func pushBytes(s *bvmCreate.chainMachine, bytes []byte) {
	s.Push(bytes)
}

func Blockhash(bench *bvmCreate.BVM) {
	pushBytes(bench.chainMachine, bench.Universe["blockHash"])
}

func Coinbase(bench *bvmCreate.BVM) {
	pushBytes(bench.chainMachine, bench.Universe["coinbase"])
}

func Timestamp(bench *bvmCreate.BVM) {
	pushBytes(bench.chainMachine, bench.Universe["timestamp"])
}

func Number(bench *bvmCreate.BVM) {
	pushBytes(bench.chainMachine, bench.Universe["number"])
}

func Difficulty(bench *bvmCreate.BVM) {
	pushBytes(bench.chainMachine, bench.Universe["difficulty"])
}

func OxygenLimit(bench *bvmCreate.BVM) {
	pushBytes(bench.chainMachine, bench.Universe["oxygenLimit"])
}
