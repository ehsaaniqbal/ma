package evaluator

import (
	"fmt"
	"os"

	"github.com/ehsaaniqbal/ma/object"
)

var builtins = map[string]*object.Builtin{
	"roar": {
		Fn: func(args ...object.Object) object.Object {
			op := ""
			for _, arg := range args {
				// Only print to stdout if not in wasm
				if os.Args[0] != "js" {
					fmt.Println(arg.Inspect())
				}
				op += fmt.Sprintf("%s\n", arg.Inspect())
			}

			return &object.String{Value: op}
		},
	},
}
