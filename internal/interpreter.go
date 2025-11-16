package internal

import(
	"os"
	"fmt"

	u "example.com/gofck/utils"
)

func Run(program *Program) {
	valid := validate(program)
	if !valid {
		fmt.Printf("Error while loading file : %s\n", program.ProgramFilePath)
		os.Exit(0)
	}
	fmt.Println("Program is valid")
}

func validate(program *Program) (bool) {
	filePathValidateResult := validateFilePath(program.ProgramFilePath)

	if filePathValidateResult {
		bracketValidateResult := validateBrackets(program.ProgramFilePath)
		if bracketValidateResult { return true }
	}

	return false
}

func validateFilePath(brainfuckFilePath string) (bool) {
	_, err := os.ReadFile(brainfuckFilePath)
	// os.Stdout.Write(data)

	if err != nil {
		fmt.Println("Error, file not found")
		return false
	}
	return true
}

func validateBrackets(brainfuckFilePath string) (bool) {
	codeBytes, _ := os.ReadFile(brainfuckFilePath)
	code := string(codeBytes)
	fmt.Println(code)
	i := 0
	var stack u.Stack
	for i < len(code) {
		if string(code[i]) == "[" {
			stack.Push(string(code[i]))
		}
		if string(code[i]) == "]" {
			stack.Pop()
		}
		i++
	}
	return stack.IsEmpty()
}
