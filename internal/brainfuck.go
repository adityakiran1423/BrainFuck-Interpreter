package internal

type Program struct {
	ProgramArray [30000]uint32
	ProgramFilePath string
	ProgramCounter int
}

func New(brainfuckFilePath string) (*Program) {
	return &Program {
		ProgramFilePath: brainfuckFilePath,
		ProgramCounter: 0,
	}
}
