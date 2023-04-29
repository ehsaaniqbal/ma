package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ehsaaniqbal/ma/evaluator"
	"github.com/ehsaaniqbal/ma/lexer"
	"github.com/ehsaaniqbal/ma/object"
	"github.com/ehsaaniqbal/ma/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "NOT MY PROBLEM MA\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
