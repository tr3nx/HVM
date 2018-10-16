package vm

// opcodes
const (
	FALSE = iota // 0
	TRUE         // 1
	ADD
	MUL
	SUB
	DIV
	GT
	LT
	EQ
	JMP
	JNZ
	JZ
	MSTORE
	MLOAD
	VLOAD
	PUSH
	POP
	PRINT
	CALL
	RET
	EXIT
)

var (
	opcodeMap = map[int]string {
		FALSE: "FALSE",
		TRUE: "TRUE",
		ADD: "ADD",
		MUL: "MUL",
		SUB: "SUB",
		DIV: "DIV",
		GT: "GT",
		LT: "LT",
		EQ: "EQ",
		JMP: "JMP",
		JNZ: "JNZ",
		JZ: "JZ",
		MSTORE: "MSTORE",
		MLOAD: "MLOAD",
		VLOAD: "VLOAD",
		PUSH: "PUSH",
		POP: "POP",
		PRINT: "PRINT",
		CALL: "CALL",
		RET: "RET",
		EXIT: "EXIT",
	}
)