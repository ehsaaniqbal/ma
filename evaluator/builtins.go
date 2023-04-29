package evaluator

import (
	"fmt"

	"github.com/ehsaaniqbal/ma/object"
)

var builtins = map[string]*object.Builtin{
	"roar": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return &object.String{Value: ""}
		},
	},
}
