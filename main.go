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

			vm.stack = append(vm.stack, v1+v2)

		case IMUL:
			v2 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v1 := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			vm.stack = append(vm.stack, v1*v2)

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

		case EXIT:
			break
		}

		if vm.debug {
			fmt.Printf("Stack: %v\n", vm.stack)
			fmt.Printf("Memory: %v\n\n", vm.memory)
		}
	}
}

func main() {
	code := []int{
		// collatz(n):
		LOAD,   // 0
		-3,     // 1
		ICONST, // 2
		2,      // 3
		IDIV,   // 4
		BRT,    // 5
		15,     // 6

		// return n / 2
		POP,    // 7
		LOAD,   // 8
		-3,     // 9
		ICONST, // 10
		2,      // 11
		IDIV,   // 12
		POP,    // 13
		RET,    // 14

		// return (3 * n) + 1
		ICONST, // 15
		3,      // 16
		LOAD,   // 17
		-3,     // 18
		IMUL,   // 19
		ICONST, // 20
		1,      // 21
		IADD,   // 22
		RET,    // 23

		// conjecture(n):
		// t := n
		LOAD,   // 24
		-3,     // 25
		MSTORE, // 26
		1,      // 27

		// while t != 1
		MLOAD,  // 28
		1,      // 29
		ICONST, // 30
		1,      // 31
		IET,    // 32

		// break
		BRT, // 33
		52,  // 34

		// t = collatz(t)
		CALL,   // 35
		0,      // 36
		1,      // 37
		MSTORE, // 38
		1,      // 39

		// nums = append(nums, t)
		// print t
		MLOAD,  // 40
		1,      // 41
		IPRINT, // 42

		// i++
		MLOAD,  // 43
		0,      // 44
		ICONST, // 45
		1,      // 46
		IADD,   // 47
		MSTORE, // 48
		0,      // 49

		JMP, // 50
		28,  // 51

		// return nums
		RET, // 52

		// main():
		ICONST, // 53
		3,      // 54
		CALL,   // 55
		24,     // 56
		1,      // 57
		EXIT,   // 58
	}

	// cpu setup
	entry := 53
	memsize := 4
	debug := true

	vm := VM{
		code:   code,
		ip:     entry,
		memory: make([]int, memsize),
		debug:  debug,
	}
	vm.Cpu()
}
