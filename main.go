package main

import "./vm"

func main() {
	code := []int{
		// fibonacci(n)
		// if n == 0: return 0
		vm.PUSH,  // 0
		0,        // 1
		vm.VLOAD, // 2
		-3,       // 3
		vm.EQ,    // 4
		vm.JNZ,   // 5
		32,       // 6

		// if n == 1: return 1
		vm.PUSH,  // 7
		1,        // 8
		vm.VLOAD, // 9
		-3,       // 10
		vm.EQ,    // 11
		vm.JNZ,   // 12
		35,       // 13

		// v1 := fib(n-1)
		vm.PUSH,  // 14
		1,        // 15
		vm.VLOAD, // 16
		-3,       // 17
		vm.SUB,   // 18
		vm.CALL,  // 19
		0,        // 20
		1,        // 21

		// v2 := fib(n-2)
		vm.PUSH,  // 22
		2,        // 23
		vm.VLOAD, // 24
		-3,       // 25
		vm.SUB,   // 26
		vm.CALL,  // 27
		0,        // 28
		1,        // 29

		// return v1 + v2
		vm.ADD, // 30
		vm.RET, // 31

		// return 0
		vm.PUSH, // 32
		0,       // 33
		vm.RET,  // 34

		// return 1
		vm.PUSH, // 35
		1,       // 36
		vm.RET,  // 37

		// main()
		// i := 0
		vm.PUSH,   // 38
		0,         // 39
		vm.MSTORE, // 40
		0,         // 41

		// while i < 10 {
		vm.PUSH,  // 42
		10,       // 43
		vm.MLOAD, // 44
		0,        // 45
		vm.LT,    // 46
		vm.JZ,    // 47
		64,       // 48

		// fibonacci(n)
		vm.MLOAD, // 49
		0,        // 50
		vm.CALL,  // 51
		0,        // 52
		1,        // 53
		vm.PRINT, // 54

		// i++
		vm.MLOAD,  // 55
		0,         // 56
		vm.PUSH,   // 57
		1,         // 58
		vm.ADD,    // 59
		vm.MSTORE, // 60
		0,         // 61

		// }
		vm.JMP, // 62
		42,     // 63

		// }
		vm.EXIT, // 64
	}

	vm.New(vm.Entry(38),
		vm.Debug(true),
		vm.RamSize(1),
		vm.Load(code)).Run()
}
