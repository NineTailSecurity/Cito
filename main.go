// main.go

package main

import (
	"fmt"
	"os"
	"os/user"

	"Cito/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to Cito Programming Language!\n", user.Username)
	fmt.Printf("Try and write a script!\n")
	repl.Start(os.Stdin, os.Stdout)
}
