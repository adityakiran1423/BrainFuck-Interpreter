package main

import(
	"fmt"
	"os"

	i "example.com/gofck/internal"
)

func main() {
	args := os.Args
	
	if len(args) != 1 {
		fmt.Println("too many arguements")
		os.Exit(0)
	} else if len(args) == 1 {
		i.Interpreter(args[1])
	} else { 
		fmt.Println("gofck") // if len(args) is 0, then run REPL? or print some info
	}
}