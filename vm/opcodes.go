package vm

// instruction opcodes
const (
	FALSE = iota
	TRUE
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
	SET
	GET
	INC
	DEC
	PRINT
	CALL
	RET
	EXIT
)

// register opcodes
const (
	A = iota
	B
	C
	D
)

var (
	registerString = []string{
		"A",
		"B",
		"C",
		"D",
	}

	opcodeString = []string{
		"FALSE",
		"TRUE",
		"ADD",
		"MUL",
		"SUB",
		"DIV",
		"GT",
		"LT",
		"EQ",
		"JMP",
		"JNZ",
		"JZ",
		"MSTORE",
		"MLOAD",
		"VLOAD",
		"PUSH",
		"POP",
		"SET",
		"GET",
		"INC",
		"DEC",
		"PRINT",
		"CALL",
		"RET",
		"EXIT",
	}
)
