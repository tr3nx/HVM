package main

import "./vm"

func main() {
	code := []int{
		vm.SET,
		vm.A,
		10,
		vm.INC,
		vm.B,
		vm.DEC,
		vm.A,
		vm.GET,
		vm.A,
		vm.PRINT,
		vm.EXIT,
	}

	vm.New(
		vm.Entry(0),
		vm.Debug(true),
		vm.RamSize(0),
		vm.Load(code),
	).Run()
}
