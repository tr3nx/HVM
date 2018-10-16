package vm

type Code struct {
	Ip int // instruction pointer
	code []int
	Current int
}

func NewCode() *Code {
	return &Code{}
}

func (c *Code) Load(code []int) {
	c.code = code
}

func (c *Code) Pop() int {
	defer func() {
		c.Ip++
	}()
	return c.Peek()
}

func (c *Code) Peek() int {
	return c.code[c.Ip]
}

func (c *Code) Length() int {
	return len(c.code)
}

func (c *Code) TranslateOp(op int) string {
	return opcodeMap[op]
}
