package internal

import (
	"log"
	"os"
)

func Interpreter(brainfuckFilePath string) {
	data, err := os.ReadFile(brainfuckFilePath)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
}

func checkBrackets() {

}