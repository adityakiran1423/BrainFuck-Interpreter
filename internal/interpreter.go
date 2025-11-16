package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	u "example.com/gofck/utils"
)

func Run(p *Program) {
	valid := validate(p)
	if !valid {
		fmt.Printf("Error while loading file : %s\n", p.ProgramFilePath)
		os.Exit(0)
	}
	fmt.Println("Program is valid")

	// actual running
	code, _ := os.ReadFile(p.ProgramFilePath)
	i := 0
	for i < len(code) {
		switch string(code[i]) {
		case "+":
			if p.ProgramArray[p.InstructionPointer] == 255 {
				p.ProgramArray[p.InstructionPointer] = 0
			} else {
				p.ProgramArray[p.InstructionPointer]++
			}

		case "-":
			if p.ProgramArray[p.InstructionPointer] == 0 {
				p.ProgramArray[p.InstructionPointer] = 255
			} else {
				p.ProgramArray[p.InstructionPointer]--
			}

		case ">":
			if p.InstructionPointer == 30000 {
				fmt.Println("Error : Memory Overflow")
				os.Exit(0)
			}
			p.InstructionPointer++

		case "<":
			if p.InstructionPointer == 0 {
				fmt.Println("Error : Memory Overflow")
				os.Exit(0)
			}
			p.InstructionPointer--

		case ".":
			fmt.Println("printing using .")
			fmt.Print(p.ProgramArray[p.InstructionPointer])

		case ",":
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter text :")
			text, _ := reader.ReadString('\n')

			if inputInt, err := strconv.Atoi(text); err == nil {
				int32Input := uint32(inputInt)
				p.ProgramArray[p.InstructionPointer] = int32Input
			} else {
				fmt.Println("Error during input")
				os.Exit(0)
			}

		case "[":


		case "]":

		default:
			fmt.Println("invalid operation")
		}
		i++
	}
	fmt.Println()
}

func validate(p *Program) bool {
	filePathValidateResult := validateFilePath(p.ProgramFilePath)

	if filePathValidateResult {
		bracketValidateResult := validateBrackets(p.ProgramFilePath)
		if bracketValidateResult {
			return true
		}
	}

	return false
}

func validateFilePath(brainfuckFilePath string) bool {
	_, err := os.ReadFile(brainfuckFilePath)

	if err != nil {
		fmt.Println("Error, file not found")
		return false
	}
	return true
}

func validateBrackets(brainfuckFilePath string) bool {
	codeBytes, _ := os.ReadFile(brainfuckFilePath)
	code := string(codeBytes)
	// fmt.Println(code)

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

	if stack.IsEmpty() {
		return true
	}

	fmt.Println("Error, brackets are not balanced")
	return false
}
