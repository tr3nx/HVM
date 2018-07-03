package main

import "fmt"

// opcode instruction set
const (
	_ = iota
	ICONST
	IADD
	IMUL
	ISUB
	IGT
	ILT
	IET
	JMP
	BRT
	BRF
	MSTORE
	MLOAD
	LOAD
	POP
	IPRINT
	CPRINT
	CALL
	RET
	HALT
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
			v1 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v2 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			vm.stack = append(vm.stack, v1+v2)

		case IMUL:
			v1 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v2 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			vm.stack = append(vm.stack, v1*v2)

		case ISUB:
			v1 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v2 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			vm.stack = append(vm.stack, v1-v2)

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

		case BRT:
			addr := vm.code[vm.ip]
			vm.ip++
			if vm.stack[len(vm.stack)-1] == TRUE {
				vm.ip = addr
			}
			vm.stack = vm.stack[:len(vm.stack)-1]

		case BRF:
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

		case HALT:
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
		// fibonacci(n)
		// if n == 0: return 0
		ICONST, // 0
		0,      // 1
		LOAD,   // 2
		-3,     // 3
		IET,    // 4
		BRT,    // 5
		32,     // 6

		// if n == 1: return 1
		ICONST, // 7
		1,      // 8
		LOAD,   // 9
		-3,     // 10
		IET,    // 11
		BRT,    // 12
		35,     // 13

		// v1 := fib(n-1)
		ICONST, // 14
		1,      // 15
		LOAD,   // 16
		-3,     // 17
		ISUB,   // 18
		CALL,   // 19
		0,      // 20
		1,      // 21

		// v2 := fib(n-2)
		ICONST, // 22
		2,      // 23
		LOAD,   // 24
		-3,     // 25
		ISUB,   // 26
		CALL,   // 27
		0,      // 28
		1,      // 29

		// return v1 + v2
		IADD, // 30
		RET,  // 31

		// return 0
		ICONST, // 32
		0,      // 33
		RET,    // 34

		// return 1
		ICONST, // 35
		1,      // 36
		RET,    // 37

		// main()
		// i := 0
		ICONST, // 38
		0,      // 39
		MSTORE, // 40
		0,      // 41

		// while i < 10 {
		ICONST, // 42
		21,     // 43
		MLOAD,  // 44
		0,      // 45
		ILT,    // 46
		BRF,    // 47
		64,     // 48

		// fibonacci(n)
		MLOAD,  // 49
		0,      // 50
		CALL,   // 51
		0,      // 52
		1,      // 53
		IPRINT, // 54

		// i++
		MLOAD,  // 55
		0,      // 56
		ICONST, // 57
		1,      // 58
		IADD,   // 59
		MSTORE, // 60
		0,      // 61

		// }
		JMP, // 62
		42,  // 63

		// }
		HALT, // 64
	}

	// cpu setup
	entry := 38
	memsize := 1
	debug := false

	vm := VM{
		code:   code,
		ip:     entry,
		memory: make([]int, memsize),
		debug:  debug,
	}
	vm.Cpu()
}
