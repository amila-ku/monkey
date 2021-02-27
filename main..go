// main.go

package main

import (
	"fmt"
	"github.com/amila-ku/monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language! \n", user.Username)
	fmt.Printf("Type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
