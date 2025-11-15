package internal

import(
	"os"
	"fmt"
)

func Run(program *Program) {
	code, valid := checkIfValidFilePath(program.ProgramFilePath)
	if !valid {
		fmt.Printf("Error while loading file : %s", program.ProgramFilePath)
	}
	// fmt.Printf("%v", code)
	_ = checkBrackets(code)
}

func checkIfValidFilePath(brainfuckFilePath string) ([]byte, bool) {
	data, err := os.ReadFile(brainfuckFilePath)
	if err != nil {
		return nil, false
	}
	os.Stdout.Write(data)
	return data, true
}

func checkBrackets(program []byte) (bool) {
	return true
}
