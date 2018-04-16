package bvmstate

import (
	"testing"

	"github.com/benchlab/bvmCreate"

	"github.com/benchlab/bvmUtils"
)

func tBlockHash(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	e := make(bvmCreate.Universe)
	bvmUtils.Assert(t, s.Size() == 0, "BVM Error! There seems to be an issue with your starting size.")
	e["blockHash"] = []byte("random")
	pushBytes(s, e["blockHash"])
}

func tCoinbase(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	e := make(bvmCreate.Universe)
	bvmUtils.Assert(t, s.Size() == 0, "BVM Error! There seems to be an issue with your starting size.")
	e["coinbase"] = []byte("random")
	pushBytes(s, e["coinbase"])
}

func tTimestamp(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	e := make(bvmCreate.Universe)
	bvmUtils.Assert(t, s.Size() == 0, "BVM Error! There seems to be an issue with your starting size.")
	e["timestamp"] = []byte("random")
	pushBytes(s, e["timestamp"])
}

func tBlockNumber(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	e := make(bvmCreate.Universe)
	bvmUtils.Assert(t, s.Size() == 0, "BVM Error! There seems to be an issue with your starting size.")
	e["blockNumber"] = []byte("random")
	pushBytes(s, e["blockNumber"])
}

func tDiff(t *testing.T) {
	s := new(bvmCreate.chainMachine)
	e := make(bvmCreate.Universe)
	bvmUtils.Assert(t, s.Size() == 0, "BVM Error! There seems to be an issue with your starting size.")
	e["difficulty"] = []byte("random")
	pushBytes(s, e["difficulty"])
}

func tOxygenLimit(t *testing.T) {

}
