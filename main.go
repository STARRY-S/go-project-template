package main

import (
	"os"

	"github.com/STARRY-S/go-project-template/commands"
)

func init() {
}

func main() {
	commands.Execute(os.Args[1:])
}
