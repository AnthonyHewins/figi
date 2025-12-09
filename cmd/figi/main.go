package main

import (
	"fmt"
	"os"
)

type commandInterface interface {
	name() string
	help() string
	run(*args) error
}

var (
	allCommands = [...]commandInterface{
		mappingCmd{},
	}
)

type args struct {
	x []string
}

func (a *args) shift() string {
	if len(a.x) == 0 {
		return ""
	}

	x := a.x[0]
	a.x = a.x[1:]
	return x
}

func main() {
	a := args{os.Args[1:]}

	cmdName := a.shift()
	var ptr commandInterface
	for _, v := range allCommands {
		switch cmdName {
		case v.name():
			ptr = v
		case "-h", "help", "--help":
			fmt.Println("todo")
			return
		}
	}

	if ptr == nil {
		fmt.Fprintln(os.Stderr, "command not found: "+cmdName)
		os.Exit(1)
	}

	if err := ptr.run(&a); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
