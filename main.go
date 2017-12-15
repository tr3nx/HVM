package main

import (
	"fmt"
)

const (
	_ = iota
	ICONST
	IADD
	IMUL
	ISUB
	IGT
	ILT
	IET
	BR
	BRT
	BRF
	MSTORE
	MLOAD
	LOAD
	POP
	PRINT
	CALL
	RET
	HALT
)

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

			vm.stack = append(vm.stack, v2-v1)

		case MSTORE:
			addr := vm.code[vm.ip]
			vm.ip++

			vm.memory[addr] = vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

		case MLOAD:
			addr := vm.code[vm.ip]
			vm.ip++

			vm.stack = append(vm.stack, vm.memory[addr])

		case PRINT:
			fmt.Println(vm.stack[len(vm.stack)-1])
			vm.stack = vm.stack[:len(vm.stack)-1]

		case IGT:
			a := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			b := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v := FALSE
			if b > a {
				v = TRUE
			}
			vm.stack = append(vm.stack, v)

		case ILT:
			a := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			b := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v := FALSE
			if b < a {
				v = TRUE
			}
			vm.stack = append(vm.stack, v)

		case IET:
			a := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			b := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]

			v := FALSE
			if b == a {
				v = TRUE
			}
			vm.stack = append(vm.stack, v)

		case BR:
			vm.ip = vm.code[vm.ip]
			vm.ip++

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

		// fmt.Printf("Stack: %v\n", vm.stack)
		// fmt.Printf("Memory: %v\n", vm.memory)
	}
}

func main() {
	fact := 0
	code := []int{
		// if N < 2 return 1
		LOAD,   // 0 - fact()
		-3,     // 1
		ICONST, // 2
		2,      // 3
		ILT,    // 4
		BRF,    // 5
		10,     // 6
		ICONST, // 7
		1,      // 8
		RET,    // 9

		// return N * fact(N-1)
		LOAD,   // 10
		-3,     // 11
		LOAD,   // 12
		-3,     // 13
		ICONST, // 14
		1,      // 15
		ISUB,   // 16
		CALL,   // 17
		fact,   // 18
		1,      // 19
		IMUL,   // 20
		RET,    // 21

		// print fact(N)
		ICONST, // 22
		5,      // 23
		MSTORE, // 24
		0,      // 25
		MLOAD,  // 26
		0,      // 27
		CALL,   // 28
		fact,   // 29
		1,      // 30
		PRINT,  // 31
		HALT,   // 32
	}

	entry := 22
	memsize := 1
	vm := VM{
		code:   code,
		ip:     entry,
		memory: make([]int, memsize),
	}
	vm.Cpu()
}
