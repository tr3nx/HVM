package vm

import "fmt"

func (vm *VM) Debug() {
	if (!vm.debugging) {
		return
	}

	opcode := vm.code.Current
	stringcode := vm.code.TranslateOp(opcode)

	fmt.Printf("Code   %v, %v\n", opcode, stringcode)
	fmt.Printf("Stack  %v\n", vm.stack.Output())
	fmt.Printf("Memory %v\n\n", vm.memory.Output())
}
