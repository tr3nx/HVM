package vm

func RamSize(size int) func(*VM) {
	return func(vm *VM) {
		vm.memory = NewMemory(size)
	}
}

func Load(code []int) func(*VM) {
	return func(vm *VM) {
		vm.code.Load(code)
	}
}

func Debug(debug bool) func(*VM) {
	return func(vm *VM) {
		vm.debugging = debug
	}
}

func Entry(entry int) func(*VM) {
	return func(vm *VM) {
		vm.code.Ip = entry
	}
}
