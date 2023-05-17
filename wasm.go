//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"syscall/js"

	"github.com/ehsaaniqbal/ma/evaluator"
	"github.com/ehsaaniqbal/ma/lexer"
	"github.com/ehsaaniqbal/ma/object"
	"github.com/ehsaaniqbal/ma/parser"
)

func eval(in string) {
	// Safely handle any runtime errors
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	env := object.NewEnvironment()
	l := lexer.New(in)
	p := parser.New(l)

	fmt.Printf("[ma]: parsing program...\n")
	program := p.ParseProgram()

	var output []map[string]string

	// Check for any parsing errors
	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			output = append(output, map[string]string{"error": fmt.Sprintf("%s\n", msg)})
		}
	} else {
		output = evaluator.EvalProgramWasm(program, env)
	}

	jsonOutput, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}

    strOutput := string(jsonOutput)
    fmt.Println("output -> ", strOutput)
	// Expose output to wasm as JSON
	js.Global().Set("ma_output", strOutput)
}

// Run evaluator on the string input sent from JS
func evalWasm(this js.Value, inputs []js.Value) interface{} {
	code := inputs[0].String()
	eval(code)
	return "[ma]: code evaluation completed"
}

func main() {
	arch := os.Args[0]
	// Check if interpreter is running on wasm
	if arch == "js" {
		// Use a channel to stop Go from killing the program early
		ch := make(chan bool)

		// Expose the evaluator function to wasm
		js.Global().Set("ma_eval", js.FuncOf(evalWasm))
		fmt.Println("[ma]: execution env set to wasm")

		<-ch
	}
}
