package vm

type Memory struct {
	memory []int
}

func NewMemory(size int) *Memory {
	return &Memory{
		memory: make([]int, size),
	}
}

func (m *Memory) Store(p, v int) {
	m.memory[p] = v
}

func (m *Memory) Get(p int) int {
	return m.memory[p]
}

func (m *Memory) Has(p int) bool {
	return m.memory[p] > 0
}

func (m *Memory) Output() interface{} {
	return m.memory
}
