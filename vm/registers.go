package vm

type Registers struct {
	Register []int
}

func NewRegisters() *Registers {
	return &Registers{
		Register: []int{
			0,
			0,
			0,
			0,
		},
	}
}

func (r *Registers) Increment(reg int) int {
	r.Register[reg]++
	return r.Register[reg]
}

func (r *Registers) Decrement(reg int) int {
	if r.Register[reg] > 0 {
		r.Register[reg]--
	}
	return r.Register[reg]
}
