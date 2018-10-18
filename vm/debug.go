package vm

import (
	"fmt"
	"strings"
)

func (vm *VM) Debug() {
	if !vm.debugging {
		return
	}

	opcode := vm.code.Current

	var rs strings.Builder
	for i, r := range vm.registers.Register {
		rs.WriteString(fmt.Sprintf("           %s: %d\n", registerString[i], r))
	}

	fmt.Printf("Code       %v, %v\n", opcode, opcodeString[opcode])
	fmt.Printf("Stack      %v\n", vm.stack.Output())
	fmt.Printf("Memory     %v\n", vm.memory.Output())
	fmt.Printf("Registers\n%v\n", rs.String())
	fmt.Printf("\n")
}
