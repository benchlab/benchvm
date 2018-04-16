package memlib

import (
	"github.com/benchlab/bvmCreate"
	"testing"
)

func TestMemory(t *testing.T) {

}

func TestNewFireMemory(t *testing.T) {

}

func TestAssertion(t *testing.T) {
	m := newMachineMemory(10)
	validMemory(m)
}

func validMemory(m bvmCreate.machineMemory) {

}
