package memlib

import (
	"fmt"
	"testing"

	"github.com/benchlab/bengo/bench"

	"github.com/benchlab/bvmUtils"
)

func tGetStorage(t *testing.T) {
	s := newMachineStorage()
	val := bengo.HashString("one")
	s.Set(bengo.HashString("hola"), val)
	bvmUtils.Assert(t, s.Get(bengo.HashString("hola")) == val, "")

}

func tCopyStorage(t *testing.T) {
	s := newMachineStorage()
	s.Set(bengo.HashString("hola"), bengo.HashString("one"))
	bvmUtils.Assert(t, s.Size() == 1, "BVM Error - Original size is incorrect..")
	f := s.Copy()
	bvmUtils.Assert(t, s.Size() == 1, "BVM Error - Original size is incorrect.")
	bvmUtils.Assert(t, f.Size() == 1, "BVM Error - Original copy size is incorrect.")
}

func tStorageSize(t *testing.T) {
	s := newMachineStorage()
	bvmUtils.Assert(t, s.Size() == 0, "BVM Error - 0 size is incorrect.")
	s.Set(bengo.HashString("hola"), bengo.HashString("one"))
	bvmUtils.Assert(t, s.Size() == 1, "BVM Error - 1 size is incorrect.")
	s.Set(bengo.HashString("hola"), bengo.HashString("one"))
	bvmUtils.Assert(t, s.Size() == 1, "BVM Error - duplicate size is incorrect.")
	s.Set(bengo.HashString("peace"), bengo.HashString("close"))
	bvmUtils.Assert(t, s.Size() == 2, fmt.Sprintf("BVM Error - 2 size is incorrect.: %d", s.Size()))
}
