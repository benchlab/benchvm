package memlib

type machineMemory struct {
	capacity    int
	store       []byte
	lastGasCost uint64
	lastReturn  []byte
}

func newMachineMemory(capacity int) *machineMemory {
	m := new(machineMemory)
	m.capacity = capacity
	return m
}

func (m *machineMemory) Size() int {
	return m.capacity
}

func (m *machineMemory) Set(offset, size uint64, value []byte) {

	if size > uint64(len(m.store)) {
		panic("INVALID memory: store empty")
	}

	if size > 0 {
		copy(m.store[offset:offset+size], value)
	}
}

func (m *machineMemory) Get(offset, size int64) (cpy []byte) {
	if size == 0 {
		return nil
	}
	if len(m.store) > int(offset) {
		cpy = make([]byte, size)
		copy(cpy, m.store[offset:offset+size])

		return
	}
	return
}

func (s machineMemory) pushViewer() {

}
