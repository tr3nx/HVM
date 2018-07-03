package main

import "fmt"

// opcode instruction set
const (
	_ = iota
	ICONST
	IADD
	IMUL
	ISUB
	IDIV
	IGT
	ILT
	IET
	JMP
	JNZ
	JZ
	MSTORE
	MLOAD
	LOAD
	POP
	IPRINT
	CPRINT
	CALL
	RET
	EXIT
)

// booleans
const (
	FALSE = iota
	TRUE
)

type VM struct {
	stack  []int
	memory []int
	code   []int
	ip     int // instruction pointer
	fp     int // frame pointer
	debug  bool
}

func (vm *VM) Cpu() {
	for vm.ip < len(vm.code) {
		opcode := vm.code[vm.ip]
		vm.ip++

		switch opcode {
		case ICONST:
			v := vm.code[vm.ip]
			vm.ip++

			vm.stack = append(vm.stack, v)

		case IADD:
			v2 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v1 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			vm.stack = append(vm.stack, v1 + v2)

		case IMUL:
			v2 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v1 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			vm.stack = append(vm.stack, v1 * v2)

		case ISUB:
			v2 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v1 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			vm.stack = append(vm.stack, v1-v2)

		case IDIV:
			v2 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v1 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			quotient := int(v1 / v2)
			remainder := int(v1 % v2)

			vm.stack = append(vm.stack, remainder)
			vm.stack = append(vm.stack, quotient)

		case MSTORE:
			addr := vm.code[vm.ip]
			vm.ip++

			vm.memory[addr] = vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

		case MLOAD:
			addr := vm.code[vm.ip]
			vm.ip++

			vm.stack = append(vm.stack, vm.memory[addr])

		case IPRINT:
			fmt.Println(vm.stack[len(vm.stack)-1])
			vm.stack = vm.stack[:len(vm.stack)-1]

		case CPRINT:
			fmt.Printf("%c", vm.stack[len(vm.stack)-1])
			vm.stack = vm.stack[:len(vm.stack)-1]

		case IGT:
			v1 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v2 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v := FALSE
			if v1 > v2 {
				v = TRUE
			}
			vm.stack = append(vm.stack, v)

		case ILT:
			v1 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v2 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v := FALSE
			if v1 < v2 {
				v = TRUE
			}
			vm.stack = append(vm.stack, v)

		case IET:
			v1 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v2 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v := FALSE
			if v1 == v2 {
				v = TRUE
			}
			vm.stack = append(vm.stack, v)

		case JMP:
			vm.ip = vm.code[vm.ip]

		case JNZ:
			addr := vm.code[vm.ip]
			vm.ip++
			if vm.stack[len(vm.stack)-1] == TRUE {
				vm.ip = addr
			}
			vm.stack = vm.stack[:len(vm.stack)-1]

		case JZ:
			addr := vm.code[vm.ip]
			vm.ip++
			if vm.stack[len(vm.stack)-1] == FALSE {
				vm.ip = addr
			}
			vm.stack = vm.stack[:len(vm.stack)-1]

		case LOAD:
			offset := vm.code[vm.ip]
			vm.ip++
			vm.stack = append(vm.stack, vm.stack[vm.fp+offset])

		case POP:
			vm.stack = vm.stack[:len(vm.stack)-1]

		case CALL:
			addr := vm.code[vm.ip]
			vm.ip++

			nargs := vm.code[vm.ip]
			vm.ip++

			// preserve frame record
			vm.stack = append(vm.stack, nargs)
			vm.stack = append(vm.stack, vm.fp)
			vm.stack = append(vm.stack, vm.ip)

			vm.fp = len(vm.stack) - 1
			vm.ip = addr

		case RET:
			retv := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			vm.ip = vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			vm.fp = vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			nargs := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			vm.stack = vm.stack[:len(vm.stack)-nargs]
			vm.stack = append(vm.stack, retv)

		case EXIT:
			break
		}

		if vm.debug {
			fmt.Printf("Stack: %v\n", vm.stack)
			fmt.Printf("Memory: %v\n", vm.memory)
		}
	}
}

func main() {
	code := []int{
		
	}

	// cpu setup
	entry := 0
	memsize := 2
	debug := false

	vm := VM{
		code:   code,
		ip:     entry,
		memory: make([]int, memsize),
		debug:  debug,
	}
	vm.Cpu()
}
