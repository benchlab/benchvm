package memlib

import "github.com/benchlab/bengo/bench"

type machineStorage struct {
	data map[bengo.Hash]bengo.Hash
}

func newMachineStorage() *machineStorage {
	s := new(machineStorage)
	s.data = make(map[bengo.Hash]bengo.Hash)
	return s
}

func (s machineStorage) Set(key, value bengo.Hash) {
	s.data[key] = value
}

func (s machineStorage) Get(key bengo.Hash) bengo.Hash {
	return s.data[key]
}

func (s machineStorage) Size() int {
	return len(s.data)
}

func (s machineStorage) Copy() (f machineStorage) {
	f.data = make(map[bengo.Hash]bengo.Hash)
	for k, v := range s.data {
		f.data[k] = v
	}
	return f
}

func (s machineStorage) pushViewer() {

}
