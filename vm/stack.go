package vm

type Stack struct {
	Fp    int // frame pointer
	stack []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v int) {
	s.stack = append(s.stack, v)
}

func (s *Stack) Pop() int {
	defer func() {
		s.stack = s.stack[:len(s.stack)-1]
	}()
	return s.stack[len(s.stack)-1]
}

func (s *Stack) Pops(argc int) []int {
	rets := make([]int, argc)
	for i := 0; i < argc; i++ {
		rets[i] = s.Pop()
	}
	return rets
}

func (s *Stack) Peek() int {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) PeekAt(i int) int {
	return s.stack[i]
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func (s *Stack) Output() interface{} {
	return s.stack
}
