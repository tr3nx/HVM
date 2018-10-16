package vm

import "fmt"

type VM struct {
	registers map[int]int
	stack  *Stack
	memory *Memory
	code   *Code
	debugging  bool
}

func New(options ...func(*VM)) *VM {
	vm := &VM{
		stack: NewStack(),
		code: NewCode(),
	}
	for _, fn := range options {
		fn(vm)
	}
	return vm
}

func (vm *VM) Run() {
	for vm.code.Ip < vm.code.Length() {
		vm.code.Current = vm.code.Pop()

		switch vm.code.Current {
		case ADD:
			v := vm.stack.Pops(2)
			vm.stack.Push(v[0] + v[1])

		case MUL:
			v := vm.stack.Pops(2)
			vm.stack.Push(v[0] * v[1])

		case SUB:
			v := vm.stack.Pops(2)
			vm.stack.Push(v[0] - v[1])

		case DIV:
			v := vm.stack.Pops(2)
			vm.stack.Push(int(v[0] % v[1])) // remainder
			vm.stack.Push(int(v[0] / v[1])) // quotient

		case GT:
			v := vm.stack.Pops(2)
			r := FALSE
			if v[0] > v[1] {
				r = TRUE
			}
			vm.stack.Push(r)

		case LT:
			v := vm.stack.Pops(2)
			r := FALSE
			if v[0] < v[1] {
				r = TRUE
			}
			vm.stack.Push(r)

		case EQ:
			v := vm.stack.Pops(2)
			r := FALSE
			if v[0] == v[1] {
				r = TRUE
			}
			vm.stack.Push(r)

		case JMP:
			vm.code.Ip = vm.code.Peek()

		case JNZ:
			addr := vm.code.Pop()
			if vm.stack.Peek() == TRUE {
				vm.code.Ip = addr
			}
			vm.stack.Pop()

		case JZ:
			addr := vm.code.Pop()
			if vm.stack.Peek() == FALSE {
				vm.code.Ip = addr
			}
			vm.stack.Pop()

		case MSTORE:
			vm.memory.Store(vm.code.Pop(), vm.stack.Pop())

		case MLOAD:
			vm.stack.Push(vm.memory.Get(vm.code.Pop()))

		case VLOAD:
			vm.stack.Push(vm.stack.PeekAt(vm.stack.Fp + vm.code.Pop()))

		case PUSH:
			vm.stack.Push(vm.code.Pop())

		case POP:
			vm.stack.Pop()

		case PRINT:
			fmt.Printf("==========\n")
			fmt.Printf("Output: %v\n", vm.stack.Pop())
			fmt.Printf("==========\n\n")

		case CALL:
			addr := vm.code.Pop()
			argc := vm.code.Pop()

			// preserve frame record
			vm.stack.Push(argc)
			vm.stack.Push(vm.stack.Fp)
			vm.stack.Push(vm.code.Ip)

			vm.stack.Fp = vm.stack.Size() - 1

			vm.code.Ip = addr

		case RET:
			v := vm.stack.Pops(4)
			retv := v[0]

			vm.code.Ip = v[1]

			vm.stack.Fp = v[2]
			argc := v[3]

			vm.stack.Pops(argc)
			vm.stack.Push(retv)

		case EXIT:
			break
		}

		vm.Debug()
	}
}
