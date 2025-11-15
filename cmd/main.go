package main

import(
	"fmt"
	"os"

	i "example.com/gofck/internal"
)

func main() {
	args := os.Args

	switch len(args) {
	case 1:
		// if len(args) is 0, then run REPL? or print some info
		fmt.Println("gofck")
		os.Exit(0)
	case 2:
		program := i.New(args[1])
		i.Run(program)
	default :
		fmt.Println("too many arguements")
		os.Exit(0)
	}
}