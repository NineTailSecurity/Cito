// main.go

package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/NineTailSecurity/Cito/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcom to Cito Programming Language!\n", user.Username)
	fmt.Printf("Try and type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
