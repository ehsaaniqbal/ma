package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/ehsaaniqbal/ma/evaluator"
	"github.com/ehsaaniqbal/ma/lexer"
	"github.com/ehsaaniqbal/ma/object"
	"github.com/ehsaaniqbal/ma/parser"
	"github.com/ehsaaniqbal/ma/repl"
)

func eval(in string) {
	env := object.NewEnvironment()
	l := lexer.New(in)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		parser.PrintParserErrors(os.Stderr, p.Errors())
		os.Exit(1)
	}

	evaluator.Eval(program, env)
}

func main() {
	args := os.Args[1:]

	// Read source code from a file
	if len(args) > 0 {
		input, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Printf("Error reading file: %s", err.Error())
		}

		eval(string(input))
	} else {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}

		// Start REPL
		fmt.Printf("Ready for placements %s ma?!\n", user.Username)
		repl.Start(os.Stdin, os.Stdout)
	}

}
