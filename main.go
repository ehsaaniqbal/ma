package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ehsaaniqbal/ma/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ready for placements %s ma?!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
