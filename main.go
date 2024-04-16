package main

import (
	"fmt"
	"github.com/WillBallentine/bark/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Bark Programming Language!\n", user.Username)
	fmt.Printf("Feel free to input some commands!\n")
	repl.Start(os.Stdin, os.Stdout)
}
